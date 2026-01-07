
# ayaneo.controller.keymap

一个 **为 AYANEO KUN / X86 掌机量身定制的 Linux 手柄按键映射工具**，使用 **Go + evdev + uinput** 实现，不依赖 X11 / Wayland，可在 **TTY / dwm / Wayland / Steam** 下工作。

> 目标：替代 antimicrox / xboxdrv，提供 **更低延迟、更可控、更可编程** 的输入映射方案。

---

## ✨ 特性

* ✅ 直接监听 `/dev/input/eventX`（内核 evdev）
* ✅ 使用 `/dev/uinput` 创建虚拟键盘
* ✅ YAML 配置，修改映射无需改代码
* ✅ 支持：

  * 手柄按键 → 键盘按键
  * 手柄按键 → 组合键（如 `Super + Enter`）
  * 摇杆 → WASD
* ✅ 不依赖桌面环境（dwm / TTY / Wayland / X11 均可）
* ✅ 适配 AYANEO KUN（XBox 360 pad / xpad 驱动）

---

## 🧩 适用设备

已验证：

* AYANEO KUN
* AYANEO 7840U / 8840U 系列（XBox pad 模式）

系统要求：

* Arch Linux（其他发行版理论可用）
* Linux kernel ≥ 6.x
* Go ≥ 1.22

---

## 📦 项目结构

```
ayanokey/
├── main.go        # 程序入口
├── config.go      # 配置加载
├── input.go       # evdev 监听
├── mapping.go     # 映射逻辑
├── uinput.go      # 虚拟键盘
├── keycode.go     # KEY_* → keycode
├── config.yaml    # 用户配置
└── README.md
```

---

## 🚀 快速开始

### 1️⃣ 安装依赖

```bash
sudo pacman -S go evtest
```

---

### 2️⃣ 确认手柄 event 设备

```bash
cat /proc/bus/input/devices
```

你应该能看到类似：

```
N: Name="Microsoft X-Box 360 pad"
H: Handlers=event6 js0
```

记住 `event6`。

---

### 3️⃣ 配置 `config.yaml`

```yaml
device: "/dev/input/event6"
deadzone: 12000

buttons:
  BTN_SOUTH: KEY_ENTER      # A
  BTN_EAST:  KEY_ESC        # B
  BTN_NORTH: KEY_TAB        # Y
  BTN_WEST:  KEY_SPACE      # X

  BTN_TL: KEY_LEFTCTRL     # LB
  BTN_TR: KEY_LEFTALT      # RB

  BTN_SELECT: KEY_BACKSPACE
  BTN_START:
    combo:
      - KEY_LEFTMETA
      - KEY_ENTER

axes:
  ABS_X:
    negative: KEY_A
    positive: KEY_D
  ABS_Y:
    negative: KEY_W
    positive: KEY_S
```

> ⚠️ 键名请以 `evtest /dev/input/eventX` 输出为准。

---

### 4️⃣ 权限设置（非常重要）

#### input 设备权限

```bash
sudo usermod -aG input $USER
```

#### uinput 权限

```bash
sudo groupadd -f uinput
sudo usermod -aG uinput $USER
```

添加 udev 规则：

```bash
sudo tee /etc/udev/rules.d/99-uinput.rules <<EOF
KERNEL=="uinput", MODE="0660", GROUP="uinput"
EOF
```

重新登录或重启。

---

### 5️⃣ 运行

```bash
go run .
```

成功启动会看到：

```
AYANEO key mapper started
Opened device: Microsoft X-Box 360 pad
```

此时按手柄，系统应收到对应的键盘输入。

---

## 🧪 调试建议

### 使用 evtest 确认按键/轴名称

```bash
sudo evtest /dev/input/event6
```

确认你在 `config.yaml` 中使用的：

* `BTN_SOUTH / BTN_EAST / ...`
* `ABS_X / ABS_Y`

与 evtest 输出完全一致。

---

## 🔒 关于 EVIOCGRAB

程序可对手柄设备进行 **独占（EVIOCGRAB）**，防止：

* Steam
* SDL
* 游戏

同时读取手柄事件。

这是 antimicrox / xboxdrv 也必须做的步骤。

---

## 🛠️ 计划中的功能

* ⏱️ Fn 键 / 背键模式切换（桌面 / 游戏）
* 🖱️ 摇杆 → 鼠标
* 🔁 热重载 config.yaml
* ⚙️ systemd user service
* 🎮 Steam 模式自动禁用

---

## 🧠 设计理念

* **内核级输入链路**（evdev → uinput）
* **最少依赖**（不绑桌面、不绑 Steam）
* **完全可控**（配置 + 代码）

如果你愿意写 Go，这是比图形化映射工具更可靠的方案。

---

## 📜 License

MIT License

---

## 🙌 致谢

* Linux evdev / uinput
* xpad 驱动
* AYANEO 社区

---

> 如果你正在使用 dwm / Arch / 掌机，这个项目就是为你准备的。

