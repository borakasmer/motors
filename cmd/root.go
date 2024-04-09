/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/borakasmer/motors/extension"
	"github.com/borakasmer/motors/parser"
	"github.com/olekukonko/tablewriter"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "motors",
	Short: "Bu Cli Tool'u ile, arabm.com'dan 5 en yeni model arac, tanımlanan filitreye göre çekilir.",
	Long: `
Tipler: 
1. Cbr650R
2. Zx6R
3. Cbr600RR
4. R25

Örnek Kullanım Şekli
----------------------
"cbr650r"
"cbr650r -t 1"
"Zx6R -t 2"
"Cbr600RR -t 3"
"R25 -t 4"

Şimdilik arabam.com'da aramada yukarıdaki modeller desteklenmektedir.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		motortype := extension.Motors.Cbr650R
		mtype, err := cmd.Flags().GetInt("type")
		if err == nil {
			switch mtype {
			case 1:
				motortype = extension.Motors.Cbr650R
			case 2:
				motortype = extension.Motors.Zx6R
			case 3:
				motortype = extension.Motors.Cbr600RR
			case 4:
				motortype = extension.Motors.R25
			default:
				motortype = extension.Motors.Cbr650R
			}
		}
		getMotorsByName(motortype)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().IntP("type", "t", 1, "Değer girilmez ise default tip CBR650R'dır.")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Cbr650R.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getMotorsByName(motorName string) {
	var resultList = parser.ParseSahibinden(motorName)
	tableHeaders := make([]string, 0)
	tableRows := make([][]string, 0)
	tableHeaders = append(tableHeaders, "Ad")
	tableHeaders = append(tableHeaders, "Baslik")
	tableHeaders = append(tableHeaders, "Yil")
	tableHeaders = append(tableHeaders, "Km")
	tableHeaders = append(tableHeaders, "Renk")
	tableHeaders = append(tableHeaders, "Fiyat")
	tableHeaders = append(tableHeaders, "Tarih")
	tableHeaders = append(tableHeaders, "Lokasyon")

	for index, result := range resultList {
		tableRow := make([]string, 0)
		tableRow = append(tableRow, result.Name.Slice())
		tableRow = append(tableRow, result.Title.Slice())
		tableRow = append(tableRow, result.Year.Slice())
		tableRow = append(tableRow, result.Km.Slice())
		tableRow = append(tableRow, result.Color.Slice())
		tableRow = append(tableRow, result.Price.Slice())
		tableRow = append(tableRow, result.CreateDate.Slice())
		tableRow = append(tableRow, result.Location.Slice())
		tableRows = append(tableRows, tableRow)
		if index == 4 {
			break
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders)
	table.SetCaption(true, "Arabam.com sitesindeki "+extension.MotorsMap[motorName]+" Son 5 aracin Fiyatları\n ®coderbora => www.borakasmer.com")
	table.AppendBulk(tableRows)

	//Set Table Color
	if !isWindows() { //Windows için Renkli Tablo başlıkları gözükmüyor...
		table.SetHeaderColor(tablewriter.Colors{
			tablewriter.Bold, tablewriter.BgMagentaColor},
			tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgGreenColor},
			tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgYellowColor},
			tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgBlueColor},
			tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgRedColor},
			tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgHiMagentaColor},
			tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgHiCyanColor},
			tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgHiBlueColor})
	}
	table.Render()
}

func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
