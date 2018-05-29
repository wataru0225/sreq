package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wasanx25/sreq/src"
)

var editor string
var lynx bool
var sort string

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search on Qiita (short-cut alias: \"s\")",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Failed to not argument of search keyword.")
			os.Exit(2)
		}

		if sort != "rel" && sort != "created" && sort != "stock" {
			fmt.Println("Please select 'rel' or 'created' or 'stock'")
			os.Exit(2)
		}

		searcher := &src.Searcher{
			Keywords:   strings.Join(args, ","),
			Pagination: 1,
			Sort:       sort,
			Lynx:       lynx,
		}
		searcher.Exec()
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&sort, "sort", "rel", "Select rel or created or stock for sort")
	searchCmd.Flags().BoolVar(&lynx, "lynx", false, "Use lynx CUI browse")
}
