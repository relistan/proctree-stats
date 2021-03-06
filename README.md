Process Tree Statistics
-----------------------

This is a single binary command line tool for printing out a pretty, colored,
process tree on Linux systems. It will also aggregate and print out stats
for Utime, Stime, Rsize, Vsize for all processes in the tree.

This is particularly useful for finding out things like "how much memory
is Nginx using?", for example.

Uncolored output example (real output is colorized):

```ansi
$ ./proctree 1

Process                                     UTime  STime   Ttl S  Ttl U  Rsize       Vsize       Ttl Rs      Ttl Vs
 \_ 1 (init)                                  981    966   150901  27140  3399680     26566656     102555648   2983481344
    \_ 1070 (getty)                             0      0       0      0  335872      16166912     335872        16166912
    \_ 1074 (getty)                             0      0       0      0  335872      16166912     335872        16166912
    \_ 1079 (getty)                             0      0       0      0  335872      16166912     335872        16166912
    \_ 1081 (getty)                             0      0       0      0  335872      16166912     335872        16166912
    \_ 1085 (getty)                             0      0       0      0  335872      16166912     335872        16166912
    \_ 1097 (cron)                             29    161      29    161  405504      19574784     405504        19574784
    \_ 1098 (atd)                               2      1       2      1  73728       17317888     73728         17317888
    \_ 1156 (dnsmasq)                           7      6       7      6  544768      26599424     544768        26599424
    \_ 1185 (login)                             2      2      12     13  745472      59985920     8916992       88375296
    |  \_ 56403 (bash)                         10     11      10     11  8171520     28389376     8171520       28389376
    \_ 1347 (vmware-vmblock-)                   0      0       0      0  438272      195522560    438272       195522560
    \_ 1366 (vmtoolsd)                      128235  11520   128235  11520  2195456     171372544    2195456      171372544
    \_ 1934 (sh)                                0      1   12427   9686  532480      4546560      11694080     462712832
    |  \_ 1973 (nginx)                          1      2   12427   9685  3678208     90968064     11161600     458166272
    |     \_ 1974 (nginx)                    3972   2732    3972   2732  1867776     91799552     1867776       91799552
    |     \_ 1975 (nginx)                    4452   3698    4452   3698  1871872     91799552     1871872       91799552
    |     \_ 1976 (nginx)                    4002   3253    4002   3253  1871872     91799552     1871872       91799552
    |     \_ 1977 (nginx)                       0      0       0      0  1871872     91799552     1871872       91799552
    \_ 27317 (docker)                        7128    745    7214    746  15228928    662433792    25174016     854372352
    |  \_ 33775 (docker)                       86      0      86      0  5386240     127066112    5386240      127066112
    |  \_ 33781 (nginx)                         0      1       0      1  2834432     32239616     4558848       64872448
    |     \_ 33816 (nginx)                      0      0       0      0  1724416     32632832     1724416       32632832
    \_ 30776 (accounts-daemon)               1042   1122    1042   1122  1204224     124452864    1204224      124452864
    \_ 31624 (bash)                             0      0       0      0  7757824     28344320     8282112       43081728
    |  \_ 31625 (nc)                            0      0       0      0  524288      14737408     524288        14737408
    \_ 36013 (rsyslogd)                       334   2051     334   2051  901120      255455232    901120       255455232
    \_ 38116 (upstart-udev-br)                356     95     356     95  622592      17645568     622592        17645568
    \_ 38118 (udevd)                          145    328     145    328  757760      21975040     1949696       65916928
    |  \_ 33729 (udevd)                         0      0       0      0  741376      21970944     741376        21970944
    |  \_ 33811 (udevd)                         0      0       0      0  450560      21970944     450560        21970944
    \_ 576 (dbus-daemon)                       14     12      14     12  892928      24518656     892928        24518656
    \_ 869 (upstart-socket-)                   14      3      14      3  483328      15556608     483328        15556608
    \_ 942 (dhclient3)                         13    165      13    165  696320      7442432      696320         7442432
    \_ 964 (sshd)                               2      6      76    265  614400      51240960     33001472     486162432
       \_ 35753 (sshd)                          0      1      64     74  3731456     79421440     21630976     247836672
       |  \_ 35767 (sshd)                       0     59      64     73  1654784     79421440     17899520     168415232
       |     \_ 35768 (bash)                   18      8      64     14  9224192     28364800     16244736      88993792
       |        \_ 35940 (vi)                  46      6      46      6  5713920     55951360     5713920       55951360
       |        \_ 35954 (proctree)             0      0       0      0  1306624     4677632      1306624        4677632
       \_ 9067 (sshd)                           0      1      10    185  1331200     79425536     10756096     187084800
          \_ 9081 (sshd)                        1    167      10    184  1179648     79425536     9424896      107659264
             \_ 9082 (bash)                     9     17       9     17  8245248     28233728     8245248       28233728

Stime:27140, Utime:150901, Rsize:25038, Vsize:2983481344
```
