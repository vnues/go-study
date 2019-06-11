package main

func mainI(){

}


// 同一个操作，作用于不同的对象，会产生不同的结果。

// work-> phone camera

/*

我们经常说：“要面向接口编程，而不是面向实现编程”。
多态性，也就要求我们面向接口编程。

不同的对象，相同的接口，但因为多态，有了不同的实现。
这样面向接口编程，就降低了耦合度，很灵活。

*/

// 怎么使用这个接口--- 多态的体现

// 一应用场景一 使用接口的方法 就是我定义方法 你给我实现这个方法就行 这种情景会比较多

// 二应用场景二 将接口当参数传入 按照duck typing谁实现这个接口也就是这个接口 -- 多态的体现