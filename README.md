# JWT Testing

Just a simple unscientific test to see how fast it takes to sign tokens

GOMAXPROCS set to 1
```
jwt-testing ❯ GOMAXPROCS=1 go run main.go
2015/04/13 11:05:55 Client #-1 took 136.492743ms
2015/04/13 11:05:55 1 request took 136.541743ms
2015/04/13 11:05:56 Client #0 took 139.799198ms
2015/04/13 11:05:56 Client #1 took 275.442165ms
2015/04/13 11:05:56 Client #2 took 421.418051ms
2015/04/13 11:05:56 Client #3 took 558.701281ms
2015/04/13 11:05:56 Client #4 took 694.418286ms
2015/04/13 11:05:56 Client #5 took 830.890032ms
2015/04/13 11:05:56 Client #6 took 966.527397ms
2015/04/13 11:05:57 Client #7 took 1.110726989s
2015/04/13 11:05:57 Client #8 took 1.251381139s
2015/04/13 11:05:57 Client #9 took 1.382923206s
2015/04/13 11:05:57 10 request serially took 1.382946359s
2015/04/13 11:05:58 Client #13 took 1.225471551s
2015/04/13 11:05:58 Client #14 took 1.23809699s
2015/04/13 11:05:58 Client #17 took 1.27958525s
2015/04/13 11:05:58 Client #19 took 1.299294913s
2015/04/13 11:05:58 Client #10 took 1.311078594s
2015/04/13 11:05:58 Client #11 took 1.318130536s
2015/04/13 11:05:58 Client #12 took 1.320444626s
2015/04/13 11:05:58 Client #15 took 1.335164147s
2015/04/13 11:05:58 Client #16 took 1.345550458s
2015/04/13 11:05:58 Client #18 took 1.352081379s
2015/04/13 11:05:58 10 request concurrently took 1.352111905s
```

GOMAXPROCS set to 16

```
jwt-testing ❯ GOMAXPROCS=16 go run main.go
2015/04/13 11:06:44 Client #-1 took 126.021889ms
2015/04/13 11:06:44 1 request took 126.077204ms
2015/04/13 11:06:44 Client #0 took 129.177184ms
2015/04/13 11:06:45 Client #1 took 263.207333ms
2015/04/13 11:06:45 Client #2 took 399.261107ms
2015/04/13 11:06:45 Client #3 took 534.970017ms
2015/04/13 11:06:45 Client #4 took 666.904214ms
2015/04/13 11:06:45 Client #5 took 799.313926ms
2015/04/13 11:06:45 Client #6 took 927.840406ms
2015/04/13 11:06:45 Client #7 took 1.054132146s
2015/04/13 11:06:45 Client #8 took 1.183236086s
2015/04/13 11:06:46 Client #9 took 1.310445368s
2015/04/13 11:06:46 10 request serially took 1.310472351s
2015/04/13 11:06:46 Client #16 took 269.916942ms
2015/04/13 11:06:46 Client #10 took 285.055123ms
2015/04/13 11:06:46 Client #15 took 288.190648ms
2015/04/13 11:06:46 Client #12 took 303.553517ms
2015/04/13 11:06:46 Client #14 took 310.845159ms
2015/04/13 11:06:46 Client #17 took 317.785958ms
2015/04/13 11:06:46 Client #18 took 325.74474ms
2015/04/13 11:06:46 Client #13 took 326.601113ms
2015/04/13 11:06:46 Client #11 took 334.499868ms
2015/04/13 11:06:46 Client #19 took 338.094607ms
2015/04/13 11:06:46 10 request concurrently took 338.121493ms
```
