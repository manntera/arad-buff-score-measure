import pyautogui
#import ctypes
import time
#import pygetwindow as gw
from PIL import ImageGrab
import pygetwindow as gw
import sys
from pywinauto.application import Application

#######################################
# 関数：指定した名前のwindowを取得する
#######################################
def activate_window(window_title):
    # ウィンドウのタイトルでウィンドウを検索する
    target_window = gw.getWindowsWithTitle(window_title)
    if target_window:
        window = target_window[0]
        # 最初のウィンドウをアクティブにする
        if window.isMinimized:
            window.restore()
        if not window.isActive:
            window.activate()
        time.sleep(1)
        return window
    else:
        print(f"ウィンドウ が見つかりませんでした。")
        sys.exit()

#######################################
# 関数：指定した名前のウィンドウ以外を最小化する
#######################################
def minimize_except_specified(title):
    all_windows = gw.getAllWindows()  # 開いているすべてのウィンドウを取得
    for window in all_windows:
        if title.lower() not in window.title.lower():  # 指定したタイトルを含まないウィンドウを検出
            window.minimize()  # ウィンドウを最小化

#######################################
### メイン処理
#######################################
print("キャプチャを開始します。マウスを触らないでお待ちください。")

#アラド戦記ウィンドウを座標取得
aradwindow = activate_window("アラド戦記")
bbox = aradwindow.left, aradwindow.top, aradwindow.left+aradwindow.width, aradwindow.top+aradwindow.height
#print("bbox:", {bbox})
            
# 「アラド戦記」という名前を含むウィンドウ以外をすべて最小化する
minimize_except_specified("アラド戦記")

#cmdウィンドウをアクティブにして左下に小さく表示
#cmdwindow = activate_window("cmd")
#cmdapp = Application().connect(handle=cmdwindow._hWnd)
#cmdwindow = cmdapp.window(handle=cmdwindow._hWnd)
#screen_width, screen_height = pyautogui.size()
#min_width, min_height = 100, 100  # 最小サイズを設定
#org_cmd_bbox = cmdwindow.left, cmdwindow.top, cmdwindow.width, cmdwindow.height
#cmdwindow.move_window(x=screen_width-min_width, y=screen_height-min_height, width=min_width, height=min_height)


# バフアイコンのスタート位置
x_coordinate = bbox[0] + round((bbox[2] - bbox[0]) * 0.1375)
y_coordinate = bbox[1] + round((bbox[3] - bbox[1]) * 0.925)

# バフアイコンの縦方向のループ
for j in range(2):
    # バフアイコンの横方向のループ
    for i in range(9):
        x_icon = x_coordinate + 21 * i
        y_icon = y_coordinate - 21 * j

        # マウスを指定した座標に移動します
        pyautogui.moveTo(x_icon, y_icon, duration=0)

        # アラド戦記ウィンドウのスクリーンショットを撮影
        screenshot = ImageGrab.grab(bbox)

        # スクリーンショットを保存
        screenshot.save(f"tmp/screenshot_{j+1}_{i+1}.png")

# 完了メッセージを表示
print("スクリーンショットを撮影しました。")

cmdwindow = activate_window("cmd")