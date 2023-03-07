package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var srcFile string
	var destAddr string
	var shorthandProgramName string = "ssgfmptupe"
	var longProgramName string = "SuperSonicGoFileMoverProTurboUltraPlusEdition"
	var version string = "0.0.1"

	// Define root command
	rootCmd := &cobra.Command{
		Use:   shorthandProgramName,
		Short: "Transfer a file from one computer to another, turbo style",
		RunE: func(cmd *cobra.Command, args []string) error {
			if srcFile == "" || destAddr == "" {
				fmt.Println(longProgramName + " " + version)
				return fmt.Errorf("src and dest options are required")
			}

			// Establish connection
			conn, err := net.Dial("tcp", destAddr)
			if err != nil {
				fmt.Println(longProgramName + " " + version)
				return fmt.Errorf("failed to connect to destination: %w", err)
			}
			defer conn.Close()

			// Transfer file
			src, err := os.Open(srcFile)
			if err != nil {
				fmt.Println(longProgramName + " " + version)
				return fmt.Errorf("failed to open source file: %w", err)
			}
			defer src.Close()

			_, err = io.Copy(conn, src)
			if err != nil {
				fmt.Println(longProgramName + " " + version)
				return fmt.Errorf("failed to transfer file: %w", err)
			}

			fmt.Println("File transfer complete")
			fmt.Println(longProgramName + " " + version)
			return nil
		},
	}

	// Define flags for root command
	rootCmd.Flags().StringVarP(&srcFile, "src", "s", "", "Path to source file")
	rootCmd.Flags().StringVarP(&destAddr, "dest", "d", "", "Destination IP address")

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
