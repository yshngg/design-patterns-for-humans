<br>
<p align="center">
  <img src="./.github/banner.svg" height="150px" />
</p>

---

[English](./README.md) | 简体中文

---

<p align="center">
🎉 设计模式的极简解释！🎉
</p>
<p align="center">
设计模式这个话题很容易让人头晕。这里我尝试用<i>最简单</i>的方式来解释它们，<br>让它们牢牢印在你的脑海里（也许也包括我的）。
</p>

---

<sub>欢迎看看我的[其他项目](http://roadmap.sh)，也可以在 [Twitter](https://twitter.com/kamrify) 上跟我打个招呼。</sub>

<br>

| [创建型设计模式 (Creational)](#创建型设计模式-creational-design-patterns) | [结构型设计模式 (Structural)](#结构型设计模式-structural-design-patterns) | [行为型设计模式 (Behavioral)](#行为型设计模式-behavioral-design-patterns) |
| :------------------------------------------------------------------------ | :------------------------------------------------------------------------ | :------------------------------------------------------------------------ |
| [简单工厂 Simple Factory](#-简单工厂-simple-factory)                      | [适配器 Adapter](#-适配器-adapter)                                        | [责任链 Chain of Responsibility](#-责任链-chain-of-responsibility)        |
| [工厂方法 Factory Method](#-工厂方法-factory-method)                      | [桥接 Bridge](#-桥接-bridge)                                              | [命令 Command](#-命令-command)                                            |
| [抽象工厂 Abstract Factory](#-抽象工厂-abstract-factory)                  | [组合 Composite](#-组合-composite)                                        | [迭代器 Iterator](#-迭代器-iterator)                                      |
| [建造者 Builder](#-建造者-builder)                                        | [装饰器 Decorator](#-装饰器-decorator)                                    | [中介者 Mediator](#-中介者-mediator)                                      |
| [原型 Prototype](#-原型-prototype)                                        | [外观 Facade](#-外观-facade)                                              | [备忘录 Memento](#-备忘录-memento)                                        |
| [单例 Singleton](#-单例-singleton)                                        | [享元 Flyweight](#-享元-flyweight)                                        | [观察者 Observer](#-观察者-observer)                                      |
|                                                                           | [代理 Proxy](#-代理-proxy)                                                | [访问者 Visitor](#-访问者-visitor)                                        |
|                                                                           |                                                                           | [策略 Strategy](#-策略-strategy)                                          |
|                                                                           |                                                                           | [状态 State](#-状态-state)                                                |
|                                                                           |                                                                           | [模板方法 Template Method](#-模板方法-template-method)                    |

<br>

简介 (Introduction)
=================

设计模式（Design Patterns）是针对反复出现的问题的解决方案；是**关于如何处理某类问题的指导原则（guidelines）**。它们不是可以直接插入应用、然后坐等奇迹发生的类、包或库。相反，它们是在特定情境下如何处理特定问题的指导原则。

> 设计模式是针对反复出现的问题的解决方案；是关于如何处理某类问题的指导原则

维基百科的描述是

> 在软件工程中，软件设计模式（software design pattern）是在软件设计中，针对特定上下文中常见问题的一种通用、可复用的解决方案。它不是一个可以直接转换成源代码或机器码的完整设计，而是一种关于如何解决问题的描述或模板，可以在许多不同的场景中使用。

⚠️ 注意
-----------------

- 设计模式不是解决所有问题的银弹（silver bullet）。
- 不要强行使用它们；如果硬套，坏事注定会发生。
- 请记住，设计模式是**针对**问题的解决方案，而不是**去寻找**问题的解决方案；所以不要想太多（overthink）。
- 如果在正确的地方以正确的方式使用，它们会成为救星；否则会导致代码一团糟。

> 另请注意，原文代码示例使用 PHP 7 编写，本人将代码实例改用更现代、简洁的 Go 编写，概念都是相通的。此外，将原文组合模式的事例换成更便于理解的文件系统。

设计模式的类型
-----------------

- [创建型 (Creational)](#创建型设计模式-creational-design-patterns)
- [结构型 (Structural)](#结构型设计模式-structural-design-patterns)
- [行为型 (Behavioral)](#行为型设计模式-behavioral-design-patterns)

创建型设计模式 (Creational Design Patterns)
==========================

简单来说

> 创建型模式关注的是如何实例化（instantiate）一个对象或一组相关对象。

维基百科的定义

> 在软件工程中，创建型设计模式是处理对象创建机制（object creation mechanisms）的设计模式，试图以适合当前情境的方式创建对象。对象创建的基本形式可能会导致设计问题或增加设计的复杂度。创建型设计模式通过以某种方式控制对象的创建来解决这个问题。

- [简单工厂 (Simple Factory)](#-简单工厂-simple-factory)
- [工厂方法 (Factory Method)](#-工厂方法-factory-method)
- [抽象工厂 (Abstract Factory)](#-抽象工厂-abstract-factory)
- [建造者 (Builder)](#-建造者-builder)
- [原型 (Prototype)](#-原型-prototype)
- [单例 (Singleton)](#-单例-singleton)

🏠 简单工厂 (Simple Factory)
--------------

现实生活中的例子

> 假设你正在盖房子，需要几扇门。你可以穿上木匠的工作服，带上木头、胶水、钉子以及造门所需的全部工具，在家里开工造门；也可以直接打电话给工厂，让他们把造好的门送过来，这样你既不需要学习任何造门的知识，也不用收拾造门带来的一堆烂摊子。

简单来说

> 简单工厂只是为客户端（client）生成一个实例，而不向客户端暴露任何实例化逻辑（instantiation logic）。

维基百科的定义

> 在面向对象编程（OOP）中，工厂（factory）是用于创建其他对象的对象——正式地说，工厂是一个函数或方法，它通过某个方法调用返回不同原型（prototype）或类（class）的对象，这个调用被假定为 "new"。

**代码示例**

首先，我们有一个门的接口（interface）及其实现（implementation）

```go
type Door interface {
	GetWidth() float32
	GetHeight() float32
}

type WoodenDoor struct {
	width  float32
	height float32
}

var _ Door = &WoodenDoor{}

func (d WoodenDoor) GetWidth() float32 {
	return d.width
}

func (d WoodenDoor) GetHeight() float32 {
	return d.height
}
```

然后，我们有一个门工厂，负责造门并返回它

```go
func MakeDoor(weight, height float32) Door {
	return WoodenDoor{
		width:  weight,
		height: height,
	}
}
```

然后可以这样使用

```go
func main() {
	// 给我造一扇 100x200 的门
	door := MakeDoor(100, 200)
	fmt.Println("Width:", door.GetWidth())
	fmt.Println("Height:", door.GetHeight())

	// 给我造一扇 50x100 的门
	door = MakeDoor(50, 100)
	fmt.Println("Width:", door.GetWidth())
	fmt.Println("Height:", door.GetHeight())
}
```

**何时使用？**

当创建一个对象不只是几条赋值语句，而是涉及一些逻辑时，把它放到一个专门的工厂中，而不是在各处重复相同的代码，是合理的做法。

🏭 工厂方法 (Factory Method)
--------------

现实生活中的例子

> 以招聘经理（hiring manager）为例。一个人不可能亲自面试所有职位。她需要根据职位空缺，决定并把面试环节委派给不同的人。

简单来说

> 它提供了一种将实例化逻辑委派（delegate）给子类（child classes）的方式。

维基百科的定义

> 在基于类的编程（class-based programming）中，工厂方法模式是一种创建型模式，它使用工厂方法来处理创建对象的问题，而无需指定所要创建对象的确切类。其做法是通过调用工厂方法来创建对象——该方法要么在接口中声明并由子类实现，要么在基类（base class）中实现并可由派生类（derived classes）选择性地重写（override）——而不是直接调用构造函数（constructor）。

**代码示例**

以上面的招聘经理为例。首先，我们有一个面试官接口及其一些实现

```go
type Interviewer interface {
	TakeInterview()
}

type Developer struct{}

var _ Interviewer = &Developer{}

func (i *Developer) TakeInterview() {
	fmt.Println("Asking about design patterns!")
}

type CommunityExecutive struct{}

var _ Interviewer = &CommunityExecutive{}

func (i *CommunityExecutive) TakeInterview() {
	fmt.Println("Asking about community building")
}
```

现在来创建我们的 `HiringManager`

```go
type HiringManager interface {
	MakeInterviewer() Interviewer
}
```

现在任何子类都可以继承（extend）它，并提供所需的面试官

```go
type DevelopmentManager struct{}

var _ HiringManager = &DevelopmentManager{}

func (dm *DevelopmentManager) MakeInterviewer() Interviewer {
	return &Developer{}
}

type MarketingManager struct{}

var _ HiringManager = &MarketingManager{}

func (mm *MarketingManager) MakeInterviewer() Interviewer {
	return &CommunityExecutive{}
}

type Title int

const (
	TitleDeveloper Title = iota
	TitleCommunityExecutive
)

func NewInterviewer(title Title) Interviewer {
	switch title {
	case TitleDeveloper:
		dm := DevelopmentManager{}
		return dm.MakeInterviewer()
	case TitleCommunityExecutive:
		mm := MarketingManager{}
		return mm.MakeInterviewer()
	}
	return nil
}
```

然后可以这样使用

```go
func main() {
	developer := NewInterviewer(TitleDeveloper)
	developer.TakeInterview() // 输出: Asking about design patterns

	communityExecutive := NewInterviewer(TitleCommunityExecutive)
	communityExecutive.TakeInterview() // 输出: Asking about community building
}
```

**何时使用？**

当类中有一些通用的处理逻辑，但所需的具体子类要在运行时（runtime）动态决定时，这个模式就很有用。换句话说，当客户端不知道自己到底需要哪个子类时。

🔨 抽象工厂 (Abstract Factory)
----------------

现实生活中的例子

> 延续简单工厂中门的例子。根据需要，你可能会从木门店买木门，从铁门店买铁门，或者从相应的店里买 PVC 门。另外，你可能还需要不同专长的师傅来安装门，比如木门需要木匠（carpenter），铁门需要焊工（welder）等等。可以看到，现在门之间存在依赖关系：木门需要木匠，铁门需要焊工，等等。

简单来说

> 工厂的工厂；一个将各自独立但相互关联/依赖的工厂组合在一起，而不指定其具体类（concrete classes）的工厂。

维基百科的定义

> 抽象工厂模式提供了一种方式，用于封装（encapsulate）一组具有共同主题的独立工厂，而无需指定它们的具体类。

**代码示例**

把上面门的例子翻译成代码。首先，我们有 `Door` 接口及其一些实现

```go
type Door interface {
	Open()
	Close()
}

type WoodenDoor struct{}

var _ Door = &WoodenDoor{}

func (w *WoodenDoor) Open() {
	fmt.Println("Opening the wooden door")
}

func (w *WoodenDoor) Close() {
	fmt.Println("Closing the wooden door")
}

type IronDoor struct{}

var _ Door = &IronDoor{}

func (w *IronDoor) Open() {
	fmt.Println("Opening the iron door")
}

func (w *IronDoor) Close() {
	fmt.Println("Closing the iron door")
}
```

然后，我们为每种门准备了相应的安装专家

```go
type DoorFittingExpert interface {
	Fit()
}

type Carpenter struct{}

var _ DoorFittingExpert = &Carpenter{}

func (c *Carpenter) Fit() {
	fmt.Println("Fitting wooden doors")
}

type Welder struct{}

var _ DoorFittingExpert = &Welder{}

func (e *Welder) Fit() {
	fmt.Println("Fitting iron doors")
}
```

现在我们有了抽象工厂，它让我们可以创建一族相关对象（family of related objects），即木门工厂会创建木门和木门安装专家，铁门工厂会创建铁门和铁门安装专家

```go
// 木门工厂，返回木匠和木门
// 木门工厂，返回木匠和木门
type DoorFactory interface {
	MakeDoor() Door
	MakeFittingExpert() DoorFittingExpert
}

type WoodenDoorFactory struct{}

func (f *WoodenDoorFactory) MakeDoor() Door {
	return &WoodenDoor{}
}

func (f *WoodenDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return &Carpenter{}
}

// 铁门工厂，返回铁门和对应的安装专家
type IronDoorFactory struct{}

func (f *IronDoorFactory) MakeDoor() Door {
	return &IronDoor{}
}

func (f *IronDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return &Welder{}
}
```

然后可以这样使用

```go
func main() {
	woodenFactory := WoodenDoorFactory{}
	door := woodenFactory.MakeDoor()
	expert := woodenFactory.MakeFittingExpert()

	expert.Fit() // 输出: Fitting wooden doors
	door.Open()  // 输出: Opening the wooden door

	ironFactory := IronDoorFactory{}
	door = ironFactory.MakeDoor()
	expert = ironFactory.MakeFittingExpert()

	expert.Fit() // 输出: Fitting iron doors
	door.Open()  // 输出: Opening the iron door
}
```

可以看到，木门工厂封装了 `carpenter` 和 `wooden door`，铁门工厂封装了 `iron door` 和 `welder`。这样就帮助我们确保：对于每一扇创建出来的门，都不会配错安装专家。

**何时使用？**

当存在相互关联的依赖关系，并且涉及不那么简单的创建逻辑时

👷 建造者 (Builder)
--------------------------------------------

现实生活中的例子

> 想象你在 Hardee's 点了一份特定的套餐，比如 "Big Hardee"，他们会*不问任何问题*直接递给你；这就是简单工厂的例子。但有些情况下，创建逻辑可能涉及更多步骤。比如你想要一份定制的 Subway 套餐，汉堡的做法有好几种选择，例如：你想要什么面包？想要哪种酱料？想要什么奶酪？等等。这种情况下，建造者模式就派上用场了。

简单来说

> 允许你创建对象的不同"口味"（flavors），同时避免构造函数污染（constructor pollution）。当一个对象可能有多种"口味"，或者创建对象涉及很多步骤时，这个模式很有用。

维基百科的定义

> 建造者模式是一种对象创建型软件设计模式，旨在为伸缩构造器反模式（telescoping constructor anti-pattern）寻找解决方案。

话虽如此，让我补充一下什么是伸缩构造器反模式。我们或多或少都见过类似下面这样的构造函数：

```python
def new(size, cheese=True, pepperoni=True, tomato=False, lettuce=True):
    pass
```

可以看到，构造函数的参数数量很快就会失控，参数的排列顺序也可能变得难以理解。而且如果将来想添加更多选项，这个参数列表还会不断增长。这就是所谓的伸缩构造器反模式。

**代码示例**

明智的替代方案是使用建造者模式。首先，我们有想要制作的汉堡

```go
type Burger struct {
	size                               int
	cheese, pepperoni, lettuce, tomato bool
}
```

然后是建造者

```go
type BurgerBuilder struct {
	size                               int
	cheese, pepperoni, lettuce, tomato bool
}

func NewBurgerBuilder(size int) *BurgerBuilder {
	return &BurgerBuilder{
		size: size,
	}
}

func (b *BurgerBuilder) AddCheese() *BurgerBuilder {
	b.cheese = true
	return b
}

func (b *BurgerBuilder) AddPepperoni() *BurgerBuilder {
	b.pepperoni = true
	return b
}

func (b *BurgerBuilder) AddLettuce() *BurgerBuilder {
	b.lettuce = true
	return b
}

func (b *BurgerBuilder) AddTomato() *BurgerBuilder {
	b.tomato = true
	return b
}

func (b *BurgerBuilder) Build() Burger {
	return Burger{
		size:      b.size,
		cheese:    b.cheese,
		pepperoni: b.pepperoni,
		lettuce:   b.lettuce,
		tomato:    b.tomato,
	}
}
```

然后可以这样使用：

```go
func main() {
	burger := NewBurgerBuilder(14).
		AddCheese().
		AddPepperoni().
		AddLettuce().
		AddTomato().
		Build()
	fmt.Println(burger) // 输出: {14 true true true true}
}
```

**何时使用？**

当一个对象可能有多种"口味"，并且需要避免伸缩构造器时使用。它与工厂模式的关键区别在于：工厂模式用于一步完成的创建过程，而建造者模式用于多步骤的创建过程。

🐑 原型 (Prototype)
------------

现实生活中的例子

> 还记得多莉（Dolly）吗？那只被克隆的羊！我们不深入细节，这里的关键点在于：一切都与克隆（cloning）有关。

简单来说

> 通过克隆现有对象来创建新对象。

维基百科的定义

> 原型模式是软件开发中的一种创建型设计模式。当要创建的对象类型由一个原型实例（prototypical instance）决定时使用该模式，通过克隆这个原型实例来产生新对象。

简而言之，它允许你创建现有对象的副本，并根据需要进行修改，而不必费力地从头创建并配置一个对象。

**代码示例**

```go
type Sheep struct {
	name     string
	category string
}

func NewSheep(name, category string) *Sheep {
	return &Sheep{
		name:     name,
		category: category,
	}
}

func (s Sheep) GetName() string {
	return s.name
}

func (s Sheep) GetCategory() string {
	return s.category
}

func (s *Sheep) SetName(name string) {
	s.name = name
}

func (s *Sheep) SetCategory(category string) {
	s.category = category
}

func (s Sheep) Clone() Sheep {
	return Sheep{
		name:     s.name,
		category: s.category,
	}
}
```

然后可以像下面这样克隆

```go
func main() {
	original := NewSheep("Jolly", "Mountain Sheep")
	fmt.Println(original.GetName())     // Jolly
	fmt.Println(original.GetCategory()) // Mountain Sheep

	cloned := original.Clone()
	cloned.SetName("Dolly")
	fmt.Println(cloned.GetName())     // Dolly
	fmt.Println(cloned.GetCategory()) // Mountain Sheep
}
```

**何时使用？**

当需要一个与现有对象相似的对象时，或者当创建对象的成本比克隆更高时。

💍 单例 (Singleton)
------------

现实生活中的例子

> 一个国家同一时间只能有一位总统。每当职责需要时，出面履职的都必须是同一位总统。这里的总统就是单例。

简单来说

> 确保某个类只会创建一个对象。

维基百科的定义

> 在软件工程中，单例模式是一种将类的实例化限制为一个对象的软件设计模式。当整个系统中恰好需要一个对象来协调各项操作时，这非常有用。

单例模式实际上被认为是一种反模式（anti-pattern），应避免过度使用。它不一定是坏的，也有一些合理的使用场景，但应谨慎使用，因为它会在应用中引入全局状态（global state），在一个地方对它的修改可能会影响到其他地方，并且调试起来会相当困难。它的另一个坏处是会让代码紧耦合（tightly coupled），而且对单例进行 mock 也可能很困难。

**代码示例**

要创建单例，需要将构造函数设为私有（private）、禁止克隆、禁止继承，并创建一个静态变量（static variable）来存放实例

```go
import (
	"sync"
)

type President struct{}

var (
	once sync.Once

	instance *President
)

func NewPresident() *President {
	once.Do(func() {
		instance = new(President)
	})

	return instance
}
```

然后这样使用

```go
func main() {
	president1 := NewPresident()
	president2 := NewPresident()
	fmt.Println(reflect.DeepEqual(president1, president2)) // true
}
```

结构型设计模式 (Structural Design Patterns)
==========================

简单来说

> 结构型模式主要关注对象组合（object composition），换句话说，就是实体之间如何相互使用。或者换一种解释：它们帮助回答"如何构建一个软件组件？"

维基百科的定义

> 在软件工程中，结构型设计模式是通过识别一种实现实体间关系的简单方式来简化设计的设计模式。

- [适配器 (Adapter)](#-适配器-adapter)
- [桥接 (Bridge)](#-桥接-bridge)
- [组合 (Composite)](#-组合-composite)
- [装饰器 (Decorator)](#-装饰器-decorator)
- [外观 (Facade)](#-外观-facade)
- [享元 (Flyweight)](#-享元-flyweight)
- [代理 (Proxy)](#-代理-proxy)

🔌 适配器 (Adapter)
-------

现实生活中的例子

> 假设你的存储卡里有一些照片，需要把它们传到电脑上。为了传输，你需要某种与电脑端口兼容的适配器，这样才能把存储卡连接到电脑上。在这个例子中，读卡器就是适配器。
> 另一个例子是大家熟知的电源适配器：三脚插头无法插入两孔插座，需要使用电源适配器使其与两孔插座兼容。
> 还有一个例子是翻译，把一个人说的话翻译给另一个人。

简单来说

> 适配器模式让你可以将一个原本不兼容的对象包装（wrap）在适配器中，使其与另一个类兼容。

维基百科的定义

> 在软件工程中，适配器模式是一种软件设计模式，它允许将现有类的接口当作另一个接口来使用。它通常用于让现有类与其他类协同工作，而无需修改其源代码。

**代码示例**

假设有一个游戏，里面有一个猎人，他会猎杀狮子。

首先，我们有一个所有类型的狮子都必须实现的 `Lion` 接口

```go
type Lion interface {
	roar()
}

type AfricanLion struct{}

var _ Lion = &AfricanLion{}

func (l *AfricanLion) roar() {}

type AsianLion struct{}

var _ Lion = &AsianLion{}

func (l *AsianLion) roar() {}
```

猎人可以猎杀任何 `Lion` 接口的实现。

```go
type Hunter struct{}

func (h *Hunter) Hunt(lion Lion) {
	lion.roar()
}
```

现在假设我们要在游戏中加入 `WildDog`（野狗），让猎人也能猎杀它。但我们不能直接这么做，因为野狗的接口不同。为了让它与猎人兼容，我们必须创建一个兼容的适配器

```go
// 需要加入游戏的新角色
type WildDog struct{}

func (d *WildDog) Bark() {}

// 包装 WildDog 的适配器，使其与游戏兼容
type WildDogAdapter struct {
	Dog WildDog
}

var _ Lion = &WildDogAdapter{}

func (d *WildDogAdapter) Roar() {
	d.Dog.Bark()
}
```

现在通过 `WildDogAdapter`，就可以在游戏中使用 `WildDog` 了。

```go
func main() {
	wildDog := WildDog{}
	wildDogAdapter := WildDogAdapter{Dog: wildDog}

	hunter := Hunter{}
	hunter.Hunt(&wildDogAdapter)
}
```

🚡 桥接 (Bridge)
------

现实生活中的例子

> 假设你有一个包含多个页面的网站，并且需要允许用户更换主题（theme）。你会怎么做？为每个主题都复制一份所有页面，还是单独创建主题，再根据用户的偏好加载？桥接模式让你可以做到后者，即：

![使用与不使用桥接模式的对比](https://cloud.githubusercontent.com/assets/11269635/23065293/33b7aea0-f515-11e6-983f-98823c9845ee.png)

简单来说

> 桥接模式讲的是优先使用组合而非继承（composition over inheritance）。实现细节从一个层次结构（hierarchy）中被推到另一个具有独立层次结构的对象中。

维基百科的定义

> 桥接模式是软件工程中使用的一种设计模式，旨在"将抽象（abstraction）与其实现（implementation）解耦，使两者可以独立变化"

**代码示例**

把上面的 WebPage 例子翻译成代码。这里是 `WebPage` 层次结构

```go
type WebPage interface {
	GetContent() string
}

type About struct {
	Theme
}

var _ WebPage = &About{}

func (a *About) GetContent() string {
	return "About page in " + a.Theme.GetColor()
}

type Careers struct {
	Theme
}

var _ WebPage = &Careers{}

func (c *Careers) GetContent() string {
	return "Careers page in " + c.Theme.GetColor()
}
```

以及独立的主题层次结构

```go
type Theme interface {
	GetColor() string
}

type DarkTheme struct{}

var _ Theme = &DarkTheme{}

func (t *DarkTheme) GetColor() string {
	return "Dark Black"
}

type LightTheme struct{}

var _ Theme = &LightTheme{}

func (t *LightTheme) GetColor() string {
	return "Off White"
}

type AquaTheme struct{}

var _ Theme = &AquaTheme{}

func (t *AquaTheme) GetColor() string {
	return "Light Blue"
}
```

然后同时使用这两个层次结构

```go
func main() {
	darkTheme := DarkTheme{}

	about := About{Theme: &darkTheme}
	careers := Careers{Theme: &darkTheme}

	fmt.Println(about.GetContent())   // About page in Dark Black
	fmt.Println(careers.GetContent()) // Careers page in Dark Black
}
```

🌿 组合 (Composite)
-----------------

现实生活中的例子

> 操作系统文件系统中有两种类型的对象，文件和文件夹。在文件系统中查找关键词时，对于文件而言，只会查看文件中的内容，对于文件夹则会在内部所有的文件或文件夹中查找关键词。

简单来说

> 组合模式让客户端以统一的方式对待各个单独的对象。

维基百科的定义

> 在软件工程中，组合模式是一种分区（partitioning）设计模式。组合模式描述的是：一组对象应当以与单个对象实例相同的方式来对待。组合的意图是将对象"组合"成树形结构（tree structures），以表示"部分-整体"层次结构（part-whole hierarchies）。实现组合模式可以让客户端统一地对待单个对象和对象组合。

**代码示例**

以上面的文件系统为例。声明统一接口，即组件接口

```go
type Component interface {
	Find(text string) bool
}
```

然后我们有文件和文件夹两种对象类型

```go

type File struct {
	name    string
	content string
}

var _ Component = &File{}

func NewFile(name, content string) *File {
	return &File{
		name:    name,
		content: content,
	}
}

func (f File) Name() string {
	return f.name
}

func (f *File) Find(text string) bool {
	return strings.Contains(f.content, text)
}

type Folder struct {
	name       string
	components []Component
}

var _ Component = &Folder{}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}

func (f Folder) Name() string {
	return f.name
}

func (f *Folder) Find(text string) bool {
	for _, component := range f.components {
		if !component.Find(text) {
			continue
		}
		return true
	}
	return false
}

func (f *Folder) AddFile(file *File) {
	if file == nil {
		return
	}
	f.components = append(f.components, file)
}

func (f *Folder) AddFolder(folder *Folder) {
	if folder == nil {
		return
	}
	f.components = append(f.components, folder)
}
```

然后可以这样使用

```go
func main() {
	file1 := NewFile("file1", "some text")
	file2 := NewFile("file2", "some text also")

	folder := NewFolder("dirname")
	folder.AddFile(file1) // 把文件放在文件夹中
	folder.AddFile(file2)

	fmt.Println(folder.Find("also")) // true
}
```

☕ 装饰器 (Decorator)
-------------

现实生活中的例子

> 想象你经营一家提供多种服务的汽车服务店。那么该如何计算账单呢？你先选定一项服务，然后动态地不断把所提供服务的价格累加上去，直到得出最终费用。这里每种服务都是一个装饰器。

简单来说

> 装饰器模式让你通过将对象包装在装饰器类的对象中，在运行时动态地改变对象的行为。

维基百科的定义

> 在面向对象编程中，装饰器模式是一种设计模式，它允许静态或动态地向单个对象添加行为，而不影响同一类中其他对象的行为。装饰器模式通常有助于遵守单一职责原则（Single Responsibility Principle），因为它允许将功能划分到关注点各不相同的类中。

**代码示例**

以咖啡为例。首先，我们有一个实现了咖啡接口的简单咖啡

```go
type Coffee interface {
	GetCost() int
	GetDescription() string
}

type SimpleCoffee struct{}

var _ Coffee = SimpleCoffee{}

func (c SimpleCoffee) GetCost() int {
	return 10
}

func (c SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}
```

我们希望代码具有可扩展性（extensible），以便在需要时可以通过选项来修改它。来做一些附加配料（add-ons，即装饰器）吧

```go
type MilkCoffee struct {
	coffee Coffee
}

var _ Coffee = MilkCoffee{}

func (c MilkCoffee) GetCost() int {
	return c.coffee.GetCost() + 2
}

func (c MilkCoffee) GetDescription() string {
	return c.coffee.GetDescription() + ", milk"
}

type WhipCoffee struct {
	coffee Coffee
}

var _ Coffee = WhipCoffee{}

func (c WhipCoffee) GetCost() int {
	return c.coffee.GetCost() + 5
}

func (c WhipCoffee) GetDescription() string {
	return c.coffee.GetDescription() + ", whip"
}

type VanillaCoffee struct {
	coffee Coffee
}

var _ Coffee = VanillaCoffee{}

func (c VanillaCoffee) GetCost() int {
	return c.coffee.GetCost() + 3
}

func (c VanillaCoffee) GetDescription() string {
	return c.coffee.GetDescription() + ", vanilla"
}
```

现在来做一杯咖啡

```go
func main() {
	var someCoffee Coffee
	someCoffee = SimpleCoffee{}
	fmt.Println(someCoffee.GetCost())        // 10
	fmt.Println(someCoffee.GetDescription()) // Simple coffee

	someCoffee = MilkCoffee{coffee: someCoffee}
	fmt.Println(someCoffee.GetCost())        // 12
	fmt.Println(someCoffee.GetDescription()) // Simple coffee, milk

	someCoffee = WhipCoffee{coffee: someCoffee}
	fmt.Println(someCoffee.GetCost())        // 17
	fmt.Println(someCoffee.GetDescription()) // Simple coffee, milk, whip

	someCoffee = VanillaCoffee{coffee: someCoffee}
	fmt.Println(someCoffee.GetCost())        // 20
	fmt.Println(someCoffee.GetDescription()) // Simple coffee, milk, whip, vanilla
}
```

📦 外观 (Facade)
----------------

现实生活中的例子

> 你是怎么开电脑的？你会说"按电源键"！你之所以这么认为，是因为你使用的是电脑对外提供的简单接口，而在内部，它必须做很多事情才能完成开机。这个面向复杂子系统（subsystem）的简单接口就是外观（门面）。

简单来说

> 外观模式为复杂的子系统提供了一个简化的接口。

维基百科的定义

> 外观是一个为更大的代码体（例如类库，class library）提供简化接口的对象。

**代码示例**

以上面的电脑为例。这里是电脑类

```go
type Computer struct{}

func (c *Computer) GetElectricShock() {
	fmt.Println("Ouch!")
}

func (c *Computer) MakeSound() {
	fmt.Println("Beep beep!")
}

func (c *Computer) ShowLoadingScreen() {
	fmt.Println("Loading..")
}

func (c *Computer) BAM() {
	fmt.Println("Ready to be used!")
}

func (c *Computer) CloseEverything() {
	fmt.Println("Bup bup bup buzzzz!")
}

func (c *Computer) Sooth() {
	fmt.Println("Zzzzz")
}

func (c *Computer) PullCurrent() {
	fmt.Println("Haaah!")
}
```

这是外观

```go
type ComputerFacade struct {
	computer *Computer
}

func (c ComputerFacade) TurnOn() {
	c.computer.GetElectricShock()
	c.computer.MakeSound()
	c.computer.ShowLoadingScreen()
	c.computer.BAM()
}

func (c ComputerFacade) TurnOff() {
	c.computer.CloseEverything()
	c.computer.PullCurrent()
	c.computer.Sooth()
}
```

现在来使用外观

```go
func main() {
	computer := ComputerFacade{computer: &Computer{}}
	computer.TurnOn()  // Ouch! Beep beep! Loading.. Ready to be used!
	computer.TurnOff() // Bup bup bup buzzzz! Haaah! Zzzzz
}
```

🍃 享元 (Flyweight)
---------

现实生活中的例子

> 你有没有在路边摊喝过现煮的茶？他们通常会比你要的多煮几杯，把剩下的留给其他顾客，以节省燃气等资源。享元模式讲的就是这个，即共享（sharing）。

简单来说

> 它通过与相似对象尽可能多地共享，来最大限度地减少内存占用或计算开销。

维基百科的定义

> 在计算机编程中，享元是一种软件设计模式。享元是一个通过与其他相似对象共享尽可能多的数据来最小化内存使用的对象；当简单的重复表示会占用不可接受的内存量时，它是一种大量使用对象的方式。

**代码示例**

把上面茶的例子翻译成代码。首先，我们有茶的种类和泡茶器

```go
// 任何会被缓存的东西都是享元 (flyweight)。
// 这里的茶的种类就是享元。
type KarakTea struct{}

type TeaMaker struct {
	availableTea map[string]*KarakTea
}

func NewTeaMaker() *TeaMaker {
	return &TeaMaker{
		availableTea: make(map[string]*KarakTea, 0),
	}
}

func (m *TeaMaker) Make(preference string) *KarakTea {
	if m.availableTea[preference] == nil {
		m.availableTea[preference] = &KarakTea{}
	}

	return m.availableTea[preference]
}
```

然后我们有 `TeaShop`（茶馆），负责接单和上茶

```go
type TeaShop struct {
	orders   map[int]*KarakTea
	teaMaker *TeaMaker
}

func NewTeaShop(teaMaker *TeaMaker) *TeaShop {
	return &TeaShop{
		orders:   make(map[int]*KarakTea, 0),
		teaMaker: teaMaker,
	}
}

func (s *TeaShop) TakeOrder(teaType string, table int) {
	s.orders[table] = s.teaMaker.Make(teaType)
}

func (s *TeaShop) Serve() {
	for table, tea := range s.orders {
		fmt.Printf("Serving tea %s to table #%d\n", tea, table)
	}
}
```

可以像下面这样使用

```go
func main() {
	teaMaker := NewTeaMaker()
	shop := NewTeaShop(teaMaker)

	shop.TakeOrder("less sugar", 1)
	shop.TakeOrder("more milk", 2)
	shop.TakeOrder("without sugar", 5)

	shop.Serve()
	// Serving tea &{} to table #1
	// Serving tea &{} to table #2
	// Serving tea &{} to table #5
}
```

🎱 代理 (Proxy)
-------------------

现实生活中的例子

> 你有没有用门禁卡开过门？打开那扇门有多种方式，可以用门禁卡打开，也可以按下一个绕过安全验证的按钮。门的主要功能是打开，但在它之上加了一层代理来增加一些功能。下面用代码示例来更好地解释一下。

简单来说

> 使用代理模式，一个类代表另一个类的功能。

维基百科的定义

> 代理在其最一般的形式中，是一个充当其他事物接口的类。代理是一个包装器（wrapper）或代理对象（agent object），客户端调用它来访问幕后真正提供服务的对象。使用代理可以只是简单地转发给真实对象，也可以提供额外的逻辑。代理中可以提供额外的功能，例如当对真实对象的操作非常消耗资源时进行缓存（caching），或者在调用真实对象的操作之前检查前置条件（preconditions）。

**代码示例**

以上面的安全门为例。首先，我们有门的接口及其一个实现

```go
type Door interface {
	Open(password string)
	Close()
}

type LabDoor struct{}

var _ Door = &LabDoor{}

func (d *LabDoor) Open(password string) {
	// Ignore the inputted password
	fmt.Println("Opening the lab door")
}

func (d *LabDoor) Close() {
	fmt.Println("Closing the lab door")
}
```

然后我们有一个代理，可以为任何我们想要的门加上安全保护

```go
type SecuredLabDoor struct {
	labDoor *LabDoor
}

var _ Door = &SecuredLabDoor{}

func NewSecuredLabDoor() *SecuredLabDoor {
	return &SecuredLabDoor{labDoor: &LabDoor{}}
}

func (d *SecuredLabDoor) Open(password string) {
	if d.authenticate(password) {
		d.labDoor.Open("")
		return
	}
	fmt.Println("Big no! It ain't possible.")
}

func (d *SecuredLabDoor) Close() {
	d.labDoor.Close()
}

func (d *SecuredLabDoor) authenticate(password string) bool {
	return password == "$ecr@t"
}
```

下面是使用方法

```go
func main() {
	securedLabDoor := NewSecuredLabDoor()
	securedLabDoor.Open("invalid") // Big no! It ain't possible.
	securedLabDoor.Open("$ecr@t")  // Opening the lab door
	securedLabDoor.Close()         // Closing the lab door
}
```

与装饰模式不同之处在于代理通常自行管理其服务对象的生命周期，而装饰的生成则总是由客户端进行控制。

另一个例子是某种数据映射器（data-mapper）的实现。例如，我最近用这个模式为 MongoDB 做了一个 ODM（Object Data Mapper，对象数据映射器），借助魔术方法 `__call()` 为 mongo 类编写了一个代理。所有方法调用都被代理到原始的 mongo 类，取回的结果原样返回；但在调用 `find` 或 `findOne` 时，数据会被映射为所需类的对象，并返回该对象而不是 `Cursor`。

行为型设计模式 (Behavioral Design Patterns)
==========================

简单来说

> 它关注的是对象之间的职责分配（assignment of responsibilities）。它们与结构型模式的不同之处在于：它们不仅指定结构，还描述了对象之间消息传递/通信的模式。换句话说，它们帮助回答"如何在软件组件中运行某个行为？"

维基百科的定义

> 在软件工程中，行为型设计模式是识别对象之间常见的通信模式并实现这些模式的设计模式。通过这样做，这些模式提高了进行这种通信的灵活性。

- [责任链 (Chain of Responsibility)](#-责任链-chain-of-responsibility)
- [命令 (Command)](#-命令-command)
- [迭代器 (Iterator)](#-迭代器-iterator)
- [中介者 (Mediator)](#-中介者-mediator)
- [备忘录 (Memento)](#-备忘录-memento)
- [观察者 (Observer)](#-观察者-observer)
- [访问者 (Visitor)](#-访问者-visitor)
- [策略 (Strategy)](#-策略-strategy)
- [状态 (State)](#-状态-state)
- [模板方法 (Template Method)](#-模板方法-template-method)

🔗 责任链 (Chain of Responsibility)
-----------------------

现实生活中的例子

> 例如，你的账户中设置了三种支付方式（`A`、`B` 和 `C`），每种方式里的金额都不同。`A` 有 100 美元，`B` 有 300 美元，`C` 有 1000 美元，支付优先级依次为 `A`、`B`、`C`。你尝试购买一件价值 210 美元的商品。使用责任链时，首先检查账户 `A` 能否完成支付，如果能，就完成支付并中断链条；如果不能，请求会继续传递到账户 `B` 检查金额，如果能则中断链条，否则请求会一直向下传递，直到找到合适的处理者（handler）。这里 `A`、`B` 和 `C` 就是链条中的环节（links），整个过程就是责任链。

简单来说

> 它帮助构建一条对象链。请求从一端进入，在对象之间不断传递，直到找到合适的处理者。

维基百科的定义

> 在面向对象设计中，责任链模式是一种由命令对象（command objects）的来源和一系列处理对象（processing objects）组成的设计模式。每个处理对象都包含定义其可处理的命令对象类型的逻辑；其余的则传递给链中的下一个处理对象。

**代码示例**

把上面账户的例子翻译成代码。首先，我们有一个包含账户链接逻辑的基础账户，以及一些具体账户

```go
type Account interface {
	Pay(amountToPay float32)
	SetNext(next Account)
}

type Bank struct {
	balance float32
	next    Account
}

var _ Account = &Bank{}

func NewBank(balance float32) *Bank {
	return &Bank{
		balance: balance,
	}
}

func (b *Bank) canPay(amount float32) bool {
	return b.balance >= amount
}

func (b *Bank) Pay(amountToPay float32) {
	if b.canPay(amountToPay) {
		b.balance -= amountToPay
		fmt.Printf("Paid %f using Bank!\n", amountToPay)
		return
	}
	fmt.Println("Cannot pay using Bank, Processing...")
	if b.next != nil {
		b.next.Pay(amountToPay)
		return
	}
	fmt.Println("Insufficient account balance.")
}

func (b *Bank) SetNext(next Account) {
	b.next = next
}

type Paypal struct {
	balance float32
	next    Account
}

var _ Account = &Paypal{}

func NewPaypal(balance float32) *Paypal {
	return &Paypal{
		balance: balance,
	}
}

func (b *Paypal) canPay(amount float32) bool {
	return b.balance >= amount
}

func (b *Paypal) Pay(amountToPay float32) {
	if b.canPay(amountToPay) {
		b.balance -= amountToPay
		fmt.Printf("Paid %f using Paypal!\n", amountToPay)
		return
	}
	fmt.Println("Cannot pay using Paypal, Processing...")
	if b.next != nil {
		b.next.Pay(amountToPay)
		return
	}
	fmt.Println("Insufficient account balance.")
}

func (b *Paypal) SetNext(next Account) {
	b.next = next
}

type Bitcoin struct {
	balance float32
	next    Account
}

var _ Account = &Bitcoin{}

func NewBitcoin(balance float32) *Bitcoin {
	return &Bitcoin{
		balance: balance,
	}
}

func (b *Bitcoin) canPay(amount float32) bool {
	return b.balance >= amount
}

func (b *Bitcoin) Pay(amountToPay float32) {
	if b.canPay(amountToPay) {
		b.balance -= amountToPay
		fmt.Printf("Paid %f using Bitcoin!\n", amountToPay)
		return
	}
	fmt.Println("Cannot pay using Bitcoin, Processing...")
	if b.next != nil {
		b.next.Pay(amountToPay)
		return
	}
	fmt.Println("Insufficient account balance.")
}

func (b *Bitcoin) SetNext(next Account) {
	b.next = next
}
```

现在用上面定义的环节（即 Bank、Paypal、Bitcoin）来准备这条链

```go
func main() {
	// 准备一条如下所示的链
	//      bank -> paypal -> bitcoin
	//
	// 第一优先级是 bank
	//      如果 bank 无法支付，则使用 paypal
	//      如果 paypal 无法支付，则使用 bitcoin

	bank := NewBank(100)       // 余额为 100 的 Bank
	paypal := NewPaypal(200)   // 余额为 200 的 Paypal
	bitcoin := NewBitcoin(300) // 余额为 300 的 Bitcoin

	bank.SetNext(paypal)
	paypal.SetNext(bitcoin)

	bank.Pay(500)

	// 输出结果
	// ==============
	// Cannot pay using Bank, Processing...
	// Cannot pay using Paypal, Processing...
	// Paid 259.000000 using Bitcoin!
}
```

👮 命令 (Command)
-------

现实生活中的例子

> 一个通用的例子是你在餐厅点餐。你（即 `Client`，客户端）让服务员（即 `Invoker`，调用者）上一些食物（即 `Command`，命令），服务员只是把请求转达给厨师（即 `Receiver`，接收者），厨师知道要做什么以及怎么做。
> 另一个例子是你（即 `Client`）用遥控器（`Invoker`）打开（即 `Command`）电视（即 `Receiver`）。

简单来说

> 允许你将操作（actions）封装在对象中。这个模式背后的关键思想是提供一种将客户端与接收者解耦（decouple）的方法。

维基百科的定义

> 在面向对象编程中，命令模式是一种行为型设计模式，它用一个对象来封装执行某个操作或在稍后触发某个事件所需的全部信息。这些信息包括方法名、拥有该方法的对象以及方法参数的值。

**代码示例**

首先，我们有接收者，它包含所有可执行操作的实现

```go
// 接收者 (Receiver)
type Bulb struct{}

func NewBulb() *Bulb {
	return &Bulb{}
}

func (b *Bulb) TurnOn() {
	fmt.Println("Bulb has been lit")
}

func (b *Bulb) TurnOff() {
	fmt.Println("Darkness!")
}
```

然后我们有一个每个命令都要实现的接口，以及一组命令

```go
// 命令 (Command)
type Command interface {
	Execute()
	Undo()
	Redo()
}

type TurnOn struct {
	bulb *Bulb
}

func NewTurnOn(bulb *Bulb) TurnOn {
	return TurnOn{bulb: bulb}
}

func (c *TurnOn) Execute() {
	c.bulb.TurnOn()
}

func (c *TurnOn) Undo() {
	c.bulb.TurnOff()
}

func (c *TurnOn) Redo() {
	c.Execute()
}

type TurnOff struct {
	bulb *Bulb
}

func NewTurnOff(bulb *Bulb) TurnOff {
	return TurnOff{bulb: bulb}
}

func (c *TurnOff) Execute() {
	c.bulb.TurnOff()
}

func (c *TurnOff) Undo() {
	c.bulb.TurnOn()
}

func (c *TurnOff) Redo() {
	c.Execute()
}
```

然后我们有一个 `Invoker`，客户端将通过它来处理各种命令

```go
// 调用者 (Invoker)
type RemoteControl struct{}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{}
}

func (rc *RemoteControl) Submit(command Command) {
	command.Execute()
}
```

最后，来看看如何在客户端中使用它

```go
func main() {
	bulb := NewBulb()
	turnOn := NewTurnOn(bulb)
	turnOff := NewTurnOff(bulb)

	remote := NewRemoteControl()
	remote.Submit(&turnOn)  // Bulb has been lit
	remote.Submit(&turnOff) // Darkness!
}
```

命令模式还可以用来实现基于事务（transaction）的系统：在执行命令的同时持续维护命令的历史记录。如果最后一个命令执行成功，一切正常；否则只需遍历历史记录，对所有已执行的命令依次执行 `undo`（撤销）。

➿ 迭代器 (Iterator)
--------

现实生活中的例子

> 老式收音机就是迭代器的一个好例子：用户可以从某个频道开始，然后使用"下一个"或"上一个"按钮来浏览各个频道。或者以 MP3 播放器或电视机为例，你可以按"下一个"和"上一个"按钮来浏览连续的频道；换句话说，它们都提供了一个用来遍历（iterate）各自频道、歌曲或电台的接口。

简单来说

> 它提供了一种访问对象中元素的方式，而不暴露其底层表示（underlying presentation）。

维基百科的定义

> 在面向对象编程中，迭代器模式是一种设计模式，它使用迭代器来遍历（traverse）容器（container）并访问容器中的元素。迭代器模式将算法与容器解耦；在某些情况下，算法必然是特定于容器的，因此无法解耦。

**代码示例**

在 Go 中，借助标准库（Standard Library）可以很容易地实现。把上面电台的例子翻译成代码。首先，我们有 `RadioStation`（电台）

```go
type RadioStation struct {
	frequency float32
}

func NewRadioStation(frequency float32) *RadioStation {
	return &RadioStation{frequency: frequency}
}

func (rs *RadioStation) Frequency() float32 {
	return rs.frequency
}
```

然后是我们的迭代器

```go
type StationList struct {
	stations []*RadioStation
}

func NewStationList() *StationList {
	return &StationList{
		stations: make([]*RadioStation, 0),
	}
}

func (sl *StationList) AddStation(station *RadioStation) {
	sl.stations = append(sl.stations, station)
}

func (sl *StationList) RemoveStation(station *RadioStation) {
	_ = slices.DeleteFunc(sl.stations, func(rs *RadioStation) bool {
		return rs.Frequency() == station.Frequency()
	})
}

func (sl *StationList) All() iter.Seq[*RadioStation] {
	return func(yield func(*RadioStation) bool) {
		for _, e := range sl.stations {
			if !yield(e) {
				return
			}
		}
	}
}
```

然后可以这样使用

```go
func main() {
	stationList := NewStationList()

	stationList.AddStation(NewRadioStation(89))
	stationList.AddStation(NewRadioStation(101))
	stationList.AddStation(NewRadioStation(102))
	stationList.AddStation(NewRadioStation(103.2))

	for rs := range stationList.All() {
		fmt.Println(rs.Frequency())
	}

	stationList.RemoveStation(NewRadioStation(89)) // 将移除频率为 89 的电台
}
```

👽 中介者 (Mediator)
========

现实生活中的例子

> 一个常见的例子是你用手机和别人通话时，你们之间有一个网络运营商（network provider），你们的通话经由它传递，而不是直接发送。在这个例子中，网络运营商就是中介者。

简单来说

> 中介者模式添加一个第三方对象（称为中介者，mediator）来控制两个对象（称为同事，colleagues）之间的交互。它有助于降低相互通信的类之间的耦合（coupling），因为现在它们不需要了解彼此的实现。

维基百科的定义

> 在软件工程中，中介者模式定义了一个对象，用于封装一组对象之间如何交互。由于它可以改变程序的运行行为，这个模式被视为一种行为型模式。

**代码示例**

下面是一个最简单的例子：一个聊天室（即中介者），以及相互发送消息的用户（即同事）。

首先，我们有中介者，即聊天室

```go
type ChatRoomMediator interface {
	ShowMessage(user *User, message string)
}

// 中介者 (Mediator)
type ChatRoom struct{}

func (r *ChatRoom) ShowMessage(user *User, message string) {
	now := time.Now()
	sender := user.Name()
	fmt.Printf("%s [%s]: %s\n", now, sender, message)
}

func NewChatRoom() ChatRoomMediator {
	return &ChatRoom{}
}
```

然后是我们的用户，即同事

```go
type User struct {
	name         string
	chatMediator ChatRoomMediator
}

func NewUser(name string, mediator ChatRoomMediator) *User {
	return &User{
		name:         name,
		chatMediator: mediator,
	}
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Send(message string) {
	u.chatMediator.ShowMessage(u, message)
}
```

以及使用方法

```go
func main() {
	mediator := NewChatRoom()
	john := NewUser("John Doe", mediator)
	jane := NewUser("Jane Doe", mediator)

	john.Send("Hi there!")
	jane.Send("Hey!")

	// 输出结果
	// 2026-11-11 17:15:38.520103 +0800 CST m=+0.000132810 [John Doe]: Hi there!
	// 2026-11-11 17:15:38.535787 +0800 CST m=+0.015816074 [Jane Doe]: Hey!
}
```

💾 备忘录 (Memento)
-------

现实生活中的例子

> 以计算器（即发起人，originator）为例：每当你执行某个计算时，上一次的计算结果会被保存在内存中（即备忘录，memento），这样你就可以回到它，并通过某些操作按钮（即负责人，caretaker）将其恢复。

简单来说

> 备忘录模式讲的是捕获并存储对象的当前状态，以便之后能够平滑地恢复。

维基百科的定义

> 备忘录模式是一种软件设计模式，它提供将对象恢复到先前状态的能力（通过回滚实现撤销，undo via rollback）。

通常在需要提供某种撤销（undo）功能时很有用。

**代码示例**

以一个文本编辑器为例，它会不时地保存状态，你可以在需要时恢复。

首先，我们有能够保存编辑器状态的备忘录对象

```go
type EditorMemento struct {
	content string
}

func NewEditorMemento(content string) *EditorMemento {
	return &EditorMemento{
		content: content,
	}
}

func (m EditorMemento) Content() string {
	return m.content
}
```

然后是我们的编辑器，即发起人，它将使用备忘录对象

```go
type Editor struct {
	content string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) Type(words string) {
	e.content += words
}

func (e *Editor) Content() string {
	return e.content
}

func (e *Editor) Save() *EditorMemento {
	return NewEditorMemento(e.content)
}

func (e *Editor) Restore(memento *EditorMemento) {
	e.content = memento.Content()
}
```

然后可以这样使用

```go
func main() {
	editor := NewEditor()

	// 输入一些内容
	editor.Type("This is the first sentence.")
	editor.Type("This is second.")

	// 保存状态，之后可恢复到: This is the first sentence. This is second.
	saved := editor.Save()

	// 再输入一些内容
	editor.Type("And this is third.")

	// 输出: 保存之前的内容
	fmt.Println(editor.Content()) // This is the first sentence. This is second. And this is third.

	editor.Restore(saved)
	fmt.Println(editor.Content()) // This is the first sentence.This is second.
}
```

😎 观察者 (Observer)
--------

现实生活中的例子

> 一个好的例子是求职者：他们订阅了某个招聘网站，每当有匹配的工作机会时就会收到通知。

简单来说

> 定义对象之间的依赖关系，使得每当一个对象改变状态时，所有依赖它的对象都会收到通知。

维基百科的定义

> 观察者模式是一种软件设计模式，其中一个称为主题（subject）的对象维护着一个依赖者列表，这些依赖者称为观察者（observers）；每当状态发生变化时，主题会自动通知它们，通常是通过调用它们的某个方法。

**代码示例**

把上面的例子翻译成代码。首先，我们有需要在职位发布时收到通知的求职者

```go
type JobPost struct {
	title string
}

func NewJobPost(title string) *JobPost {
	return &JobPost{
		title: title,
	}
}

func (jp *JobPost) Title() string {
	return jp.title
}

type Observer interface {
	OnJobPosted(job *JobPost)
}


type JobSeeker struct {
	name string
}

var _ Observer = &JobSeeker{}

func NewJobSeeker(name string) *JobSeeker {
	return &JobSeeker{
		name: name,
	}
}

func (js *JobSeeker) OnJobPosted(job *JobPost) {
	fmt.Printf("Hi %s! New job posted: %s\n", js.name, job.Title())
}
```

然后是求职者将要订阅的职位发布方

```go
type Observable interface {
	Notify(jobPosting *JobPost)
	Attach(observer Observer)
	AddJob(jobPosting *JobPost)
}

type EmploymentAgency struct {
	observers []Observer
}

var _ Observable = &EmploymentAgency{}

func NewEmploymentAgency() *EmploymentAgency {
	return &EmploymentAgency{
		observers: make([]Observer, 0),
	}
}

func (ea *EmploymentAgency) Notify(jobPosting *JobPost) {
	for _, observer := range ea.observers {
		observer.OnJobPosted(jobPosting)
	}
}

func (ea *EmploymentAgency) Attach(observer Observer) {
	ea.observers = append(ea.observers, observer)
}

func (ea *EmploymentAgency) AddJob(jobPosting *JobPost) {
	ea.Notify(jobPosting)
}
```

然后可以这样使用

```go
func main() {
  // 创建订阅者 (subscribers)
	johnDoe := NewJobSeeker("John Doe")
	janeDoe := NewJobSeeker("Jane Doe")

  // 创建发布者 (publisher) 并添加订阅者
	jobPostings := NewEmploymentAgency()
	jobPostings.Attach(johnDoe)
	jobPostings.Attach(janeDoe)

  // 添加一个新职位，看看订阅者是否收到通知
	jobPostings.AddJob(NewJobPost("Software Engineer"))
}


// 输出
// Hi John Doe! New job posted: Software Engineer
// Hi Jane Doe! New job posted: Software Engineer
```

🏃 访问者 (Visitor)
-------

现实生活中的例子

> 假设有人要去迪拜旅游。他们只需要一种方式（即签证）进入迪拜。抵达之后，他们就可以自行游览迪拜的任何地方，无需申请许可，也不用为了游览某处而四处奔波；只要告诉他们一个地方，他们就能去游览。访问者模式让你可以做到这一点：它帮助你添加可供游览的地方，让他们尽可能多地游览，而无需做任何跑腿工作。

简单来说

> 访问者模式让你可以为对象添加更多操作，而无需修改这些对象。

维基百科的定义

> 在面向对象编程和软件工程中，访问者设计模式是一种将算法与其所操作的对象结构分离的方式。这种分离的一个实际效果是：可以在不修改现有对象结构的情况下为其添加新操作。这是遵循开闭原则（open/closed principle）的一种方式。

**代码示例**

以一个动物园模拟为例：里面有几种不同的动物，我们需要让它们发出声音。下面用访问者模式来实现

```go
// 被访问者 (Visitee)
type Animal interface {
	Accept(operation AnimalOperation)
}

// 访问者 (Visitor)
type AnimalOperation interface {
	VisitMonkey(monkey *Monkey)
	VisitLion(lion *Lion)
	VisitDolphin(dolphin *Dolphin)
}
```

然后是动物的实现

```go
type Monkey struct{}

func (m *Monkey) Shout() {
	fmt.Println("Ooh oo aa aa!")
}

func (m *Monkey) Accept(operation AnimalOperation) {
	operation.VisitMonkey(m)
}

type Lion struct{}

func (l *Lion) Roar() {
	fmt.Println("Roaaar!")
}

func (l *Lion) Accept(operation AnimalOperation) {
	operation.VisitLion(l)
}

type Dolphin struct{}

func (d *Dolphin) Speak() {
	fmt.Println("Tuut tuttu tuutt!")
}

func (d *Dolphin) Accept(operation AnimalOperation) {
	operation.VisitDolphin(d)
}
```

来实现我们的访问者

```go
type Speak struct{}

var _ AnimalOperation = &Speak{}

func (s *Speak) VisitMonkey(monkey *Monkey) {
	monkey.Shout()
}

func (s *Speak) VisitLion(lion *Lion) {
	lion.Roar()
}

func (s *Speak) VisitDolphin(dolphin *Dolphin) {
	dolphin.Speak()
}
```

然后可以这样使用

```go
func main() {
	monkey := &Monkey{}
	lion := &Lion{}
	dolphin := &Dolphin{}

	speak := &Speak{}
	monkey.Accept(speak)  // Ooh oo aa aa!
	lion.Accept(speak)    // Roaaar!
	dolphin.Accept(speak) // Tuut tuttu tuutt!
}
```

我们本可以简单地为动物建立一个继承层次结构（inheritance hierarchy）来实现这一点，但那样每当需要为动物添加新动作时，就必须修改动物类。而现在我们不需要修改它们。例如，假设我们被要求为动物添加跳跃行为，只需创建一个新的访问者即可，即

```go
type Jump struct{}

var _ AnimalOperation = &Jump{}

func (j *Jump) VisitMonkey(monkey *Monkey) {
	fmt.Println("Jumped 20 feet high! on to the tree!")
}

func (s *Jump) VisitLion(lion *Lion) {
	fmt.Println("Jumped 7 feet! Back on the ground!")
}

func (s *Jump) VisitDolphin(dolphin *Dolphin) {
	fmt.Println("Walked on water a little and disappeared")
}
```

以及使用方法

```go
func main() {
	monkey := &Monkey{}
	lion := &Lion{}
	dolphin := &Dolphin{}

	jump := &Jump{}
	monkey.Accept(jump)  // Jumped 20 feet high! on to the tree!
	lion.Accept(jump)    // Jumped 7 feet! Back on the ground!
	dolphin.Accept(jump) // Walked on water a little and disappeared
}
```

💡 策略 (Strategy)
--------

现实生活中的例子

> 以排序为例：我们实现了冒泡排序（bubble sort），但数据开始增长，冒泡排序变得非常慢。为了解决这个问题，我们实现了快速排序（quick sort）。但现在，虽然快速排序算法在大数据集上表现更好，在小数据集上却非常慢。为了处理这种情况，我们实现了一个策略：小数据集使用冒泡排序，大数据集使用快速排序。

简单来说

> 策略模式允许你根据情况切换算法或策略。

维基百科的定义

> 在计算机编程中，策略模式（也称为政策模式，policy pattern）是一种行为型软件设计模式，它允许在运行时选择算法的行为。

**代码示例**

把上面的例子翻译成代码。首先，我们有策略接口和不同的策略实现

```go
type SortStrategy interface {
	Sort(array []int) []int
}

type BubbleSortStrategy struct{}

var _ SortStrategy = &BubbleSortStrategy{}

func (bs *BubbleSortStrategy) Sort(array []int) []int {
	fmt.Println("Sorting using bubble sort")
	// 执行排序
	return array
}

type QuickSortStrategy struct{}

var _ SortStrategy = &QuickSortStrategy{}

func (qs *QuickSortStrategy) Sort(array []int) []int {
	fmt.Println("Sorting using quick sort")
	// 执行排序
	return array
}
```

然后是将要使用任意策略的客户端

```go
type Sorter struct {
	sorterSmall SortStrategy
	sorterBig   SortStrategy
}

func NewSorter(sorterSmall, sorterBig SortStrategy) *Sorter {
	return &Sorter{
		sorterSmall: sorterSmall,
		sorterBig:   sorterBig,
	}
}

func (s *Sorter) Sort(array []int) []int {
	if len(array) > 5 {
		return s.sorterBig.Sort(array)
	}
	return s.sorterSmall.Sort(array)
}
```

可以这样使用

```go
func main() {
	small := []int{1, 3, 4, 2}
	big := []int{1, 4, 3, 2, 8, 10, 5, 6, 9, 7}

	sorter := NewSorter(&BubbleSortStrategy{}, &QuickSortStrategy{})

	sorter.Sort(small) // 输出 : Sorting using bubble sort
	sorter.Sort(big)   // 输出 : Sorting using quick sort
}
```

💢 状态 (State)
-----

现实生活中的例子

> 想象你正在使用某个绘图应用，选择了画笔来绘画。画笔会根据所选颜色改变其行为：如果你选择了红色，它就画红色；如果选择了蓝色，它就画蓝色，等等。

简单来说

> 它让你在状态改变时改变类的行为。

维基百科的定义

> 状态模式是一种行为型软件设计模式，它以面向对象的方式实现状态机（state machine）。使用状态模式时，状态机的实现方式是：将每个单独的状态实现为状态模式接口的派生类，并通过调用模式超类（superclass）中定义的方法来实现状态转换（state transitions）。
> 状态模式可以被理解为一种策略模式，它能够通过调用模式接口中定义的方法来切换当前策略。

**代码示例**

以电话为例。首先，我们有状态接口和一些状态实现

```go
type PhoneState interface {
	PickUp() (PhoneState, error)
	HangUp() (PhoneState, error)
	Dial() (PhoneState, error)
}

type PhoneStateIdle struct{}

var _ PhoneState = &PhoneStateIdle{}

func (ps *PhoneStateIdle) PickUp() (PhoneState, error) {
	return &PhoneStatePickedUp{}, nil
}

func (ps *PhoneStateIdle) HangUp() (PhoneState, error) {
	return nil, errors.New("already idle")
}

func (ps *PhoneStateIdle) Dial() (PhoneState, error) {
	return nil, errors.New("unable to dial in idle state")
}

type PhoneStatePickedUp struct{}

var _ PhoneState = &PhoneStatePickedUp{}

func (ps *PhoneStatePickedUp) PickUp() (PhoneState, error) {
	return nil, errors.New("already picked up")
}

func (ps *PhoneStatePickedUp) HangUp() (PhoneState, error) {
	return &PhoneStateIdle{}, nil
}

func (ps *PhoneStatePickedUp) Dial() (PhoneState, error) {
	return &PhoneStateCalling{}, nil
}

type PhoneStateCalling struct{}

var _ PhoneState = &PhoneStateCalling{}

func (ps *PhoneStateCalling) PickUp() (PhoneState, error) {
	return nil, errors.New("already picked up")
}

func (ps *PhoneStateCalling) HangUp() (PhoneState, error) {
	return &PhoneStateIdle{}, nil
}

func (ps *PhoneStateCalling) Dial() (PhoneState, error) {
	return nil, errors.New("already dialing")
}
```

然后是我们的 Phone 类，它会在不同的行为调用时改变状态

```go
type Phone struct {
	state PhoneState
}

func NewPhone() *Phone {
	return &Phone{
		state: &PhoneStateIdle{},
	}
}

func (p *Phone) PickUp() (err error) {
	p.state, err = p.state.PickUp()
	return
}

func (p *Phone) HangUP() (err error) {
	p.state, err = p.state.HangUp()
	return
}

func (p *Phone) Dial() (err error) {
	p.state, err = p.state.Dial()
	return
}
```

然后可以如下使用，它会调用相应状态的方法：

```go
func main() {
	phone := NewPhone()

	phone.PickUp()
	phone.Dial()
}
```

📒 模板方法 (Template Method)
---------------

现实生活中的例子

> 假设我们要盖一栋房子。建造步骤可能如下：
>
> - 打好房子的地基
> - 砌墙
> - 加盖屋顶
> - 加盖其他楼层

> 这些步骤的顺序永远不能改变，比如你不能在砌墙之前盖屋顶等等，但每个步骤都可以修改，例如墙可以用木头、聚酯或石头来建。

简单来说

> 模板方法定义了某个算法执行方式的骨架（skeleton），但将这些步骤的实现推迟（defer）到子类中。

维基百科的定义

> 在软件工程中，模板方法模式是一种行为型设计模式，它在一个操作中定义算法的程序骨架，将某些步骤推迟到子类中实现。它让人可以在不改变算法结构的情况下重新定义算法的某些步骤。

**代码示例**

假设我们有一个构建工具（build tool），帮助我们测试（test）、代码检查（lint）、构建（build）、生成构建报告（即代码覆盖率报告、lint 报告等），并将应用部署（deploy）到测试服务器上。

首先，我们有一个基类，它指定了构建算法的骨架

```go
type Builder interface {
	test()
	lint()
	assemble()
	deploy()
}

type builder struct {
	Builder
}

func (b *builder) build() {
	b.test()
	b.lint()
	b.assemble()
	b.deploy()
}
```

然后是具体的实现

```go
type AndroidBuilder struct{}

var _ Builder = &AndroidBuilder{}

func (ab *AndroidBuilder) test() {
	fmt.Println("Running android tests")
}

func (ab *AndroidBuilder) lint() {
	fmt.Println("Linting the android code")
}

func (ab *AndroidBuilder) assemble() {
	fmt.Println("Assembling the android build")
}

func (ab *AndroidBuilder) deploy() {
	fmt.Println("Deploying android build to server")
}

type IosBuilder struct{}

var _ Builder = &IosBuilder{}

func (ab *IosBuilder) test() {
	fmt.Println("Running ios tests")
}

func (ab *IosBuilder) lint() {
	fmt.Println("Linting the ios code")
}

func (ab *IosBuilder) assemble() {
	fmt.Println("Assembling the ios build")
}

func (ab *IosBuilder) deploy() {
	fmt.Println("Deploying ios build to server")
}
```

然后可以这样使用

```go
func main() {
	androidBuilder := &AndroidBuilder{}
	b := builder{
		Builder: androidBuilder,
	}
	b.build()
	// 输出:
	// Running android tests
	// Linting the android code
	// Assembling the android build
	// Deploying android build to server

	iosBuilder := &IosBuilder{}
	b = builder{
		Builder: iosBuilder,
	}
	b.build()
	// 输出:
	// Running ios tests
	// Linting the ios code
	// Assembling the ios build
	// Deploying ios build to server
}
```

## 🚦 总结 (Wrap Up)

差不多就这些了。我会继续完善这份内容，所以你可以 watch/star 这个仓库，方便以后回来查看。另外，我还计划写一份关于架构模式（architectural patterns）的同类内容，敬请期待。

## 👬 贡献 (Contribution)

- 报告问题（issues）
- 提交改进的 pull request
- 帮忙宣传推广
- 欢迎反馈任何意见 [![Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/kamrify.svg?style=social&label=Follow%20%40kamrify)](https://twitter.com/kamrify)

## 许可证 (License)

[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)
