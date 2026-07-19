package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"termcode/internal/application/provider"
)

var providerListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List configured providers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runProviderList()
	},
}

var providerDeleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Aliases: []string{"rm", "remove"},
	Short:   "Delete a provider by ID",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runProviderDelete(args[0])
	},
}

func init() {
	providerCmd.AddCommand(providerListCmd)
	providerCmd.AddCommand(providerDeleteCmd)
}

func runProviderList() error {
	svc, closeDB, err := openProviderService()
	if err != nil {
		return err
	}
	defer closeDB()

	providers, err := svc.List(context.Background())
	if err != nil {
		return fmt.Errorf("list providers: %w", err)
	}

	if len(providers) == 0 {
		fmt.Println("No providers configured. Use 'tc chat' to add one.")
		return nil
	}

	fmt.Printf("\n  %-4s %-20s %-30s %-14s %s\n", "NUM", "NAME", "BASE URL", "STATUS", "DEFAULT")
	fmt.Printf("  %s\n", strings.Repeat("-", 80))

	for i, p := range providers {
		def := ""
		if p.IsDefault {
			def = " (*)"
		}
		fmt.Printf("  %-4d %-20s %-30s %-14s %s\n",
			i+1, p.Name, p.BaseURL, p.Status, def)
	}

	fmt.Println()
	fmt.Print("Enter number to delete (or 0 to cancel): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil || num < 0 || num > len(providers) {
		fmt.Println("Cancelled.")
		return nil
	}

	if num == 0 {
		fmt.Println("Cancelled.")
		return nil
	}

	target := providers[num-1]
	fmt.Printf("Delete provider %q (%s)? [y/N]: ", target.Name, target.ID)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("Cancelled.")
		return nil
	}

	if err := svc.Delete(context.Background(), target.ID); err != nil {
		return fmt.Errorf("delete provider: %w", err)
	}

	fmt.Printf("Deleted provider %q.\n", target.Name)
	return nil
}

func runProviderDelete(id string) error {
	svc, closeDB, err := openProviderService()
	if err != nil {
		return err
	}
	defer closeDB()

	p, err := svc.GetByID(context.Background(), id)
	if err != nil {
		return fmt.Errorf("find provider %q: %w", id, err)
	}

	fmt.Printf("Delete provider %q (%s)? [y/N]: ", p.Name, p.ID)
	reader := bufio.NewReader(os.Stdin)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("Cancelled.")
		return nil
	}

	if err := svc.Delete(context.Background(), id); err != nil {
		return fmt.Errorf("delete provider: %w", err)
	}

	fmt.Printf("Deleted provider %q.\n", p.Name)
	return nil
}

func openProviderService() (*provider.Service, func(), error) {
	db, err := openDB()
	if err != nil {
		return nil, nil, err
	}

	if err := runMigrations(db); err != nil {
		db.Close()
		return nil, nil, err
	}

	repo := openProviderRepo(db)
	svc := provider.NewService(repo, logger)
	return svc, func() { db.Close() }, nil
}
