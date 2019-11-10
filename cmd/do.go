package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trungtvq/craw/client/craw"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do example",
	Run:   doRun,
}

func doRun(cmd *cobra.Command, args []string) {
	fmt.Println(os.Getenv("ONEDRIVE"))
	fmt.Println(craw.GetLink("https://www.fshare.vn/file/3WZ2XAERFLDUVM7?token=1573397899"))
	//fmt.Println(craw.Login("truongvqtrung@gmail.com", "T0id4ym4h0i!"))
	craw.GetLinkFromFile("./f")
}
