package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    "image/color"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Proxy Agent UI")

    // Sidebar menu
    sidebar := container.NewVBox(
        canvas.NewText("用户深蓝", color.White),
        widget.NewLabel("剩余 157 天"),
        widget.NewButton("会员充值", func() {}),
        widget.NewButtonWithIcon("我的游戏", theme.MediaPlayIcon(), func() {}),
        widget.NewButtonWithIcon("全部游戏", theme.ListIcon(), func() {}),
    )

    // Game acceleration cards
    steamCard := container.NewVBox(
        canvas.NewImageFromFile("steam.png"), // Placeholder for Steam image
        widget.NewLabel("Steam商店"),
        widget.NewButton("一键加速", func() {}),
    )
    wowCard := container.NewVBox(
        canvas.NewImageFromFile("wow.png"), // Placeholder for WoW image
        widget.NewLabel("魔兽世界国际服"),
        widget.NewButton("一键加速", func() {}),
    )
    diabloCard := container.NewVBox(
        canvas.NewImageFromFile("diablo.png"), // Placeholder for Diablo image
        widget.NewLabel("PS-暗黑破坏神IV"),
        widget.NewButton("一键加速", func() {}),
    )

    gamesSection := container.NewGridWithColumns(3, steamCard, wowCard, diabloCard)

    // Bottom banner
    bottomBanner := widget.NewButton("全新双千兆加速盒2 - 立即抢购", func() {})

    // Combine everything into a layout
    content := container.NewBorder(
        nil, bottomBanner, sidebar, nil, gamesSection,
    )

    myWindow.SetContent(content)
    myWindow.Resize(fyne.NewSize(800, 600))
    myWindow.ShowAndRun()
}
