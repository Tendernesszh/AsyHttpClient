## Reactive动机<br>
首先，Reactive动机的作用就是将服务端的需要被客户端调用内部服务操作公开到<br>
编排层，从而避免服务端所有内部服务暴露在外部世界。此外，Reactive动机还使<br>
得流量开销减少以及延迟降低。<br>
![image](https://github.com/Tendernesszh/AsyHttpClient/blob/master/AsyHttpClient/6-1.png)<br>
## 使用同步方法解决问题<br>
![image](https://github.com/Tendernesszh/AsyHttpClient/blob/master/AsyHttpClient/6-2.png)<br>
## 搭建异步机制解决问题<br>
![image](https://github.com/Tendernesszh/AsyHttpClient/blob/master/AsyHttpClient/6-3.png)<br>
## 测试结果<br>
![image](https://github.com/Tendernesszh/AsyHttpClient/blob/master/AsyHttpClient/%E6%B5%8B%E8%AF%95%E7%BB%93%E6%9E%9C.png)<br>
可以发现，使用异步机制的速度会比使用同步的速度快很多。
