## example

### Run json.go

```bash
&{example.json 1 0x212650 {{0 0} 0 0 0 0} map[a:Easy! b:map[c:map[e:Just Do it f:2 g:ON t:1day] d:[3 4]] h:1.01]}
get a Easy!				= Easy!
get a.b.c def:example	= example
get b.c.e				= Just Do it
get h is 1.01			= 1.01
get b.c.f def:100,but 2	= 2
get b.c.e not exist		= 0
get b.c.g ON,return T	= true
get b.c.x def:true		= true
get b config			= &{ 0 0x212650 {{0 0} 0 0 0 0} map[b:map[c:map[f:2 g:ON t:1day e:Just Do it] d:[3 4]]]}
get b.d list 3->4		= [3 4]
get b.c.t time:1day		= 24h0m0s
set a.b.c Correct		= <nil>
set b.c.e Correct		= <nil>
set b.c.d d				= <nil>
set b.c.g false 		= <nil>
set b.d list 1->4		= <nil>
get a def:example		= example
get a interface			= map[b:map[c:Correct]]
get a.b.c set Correct	= Correct
get b.c.e set Correct	= Correct
get b.c.g set false		= false
get b.c.d set d			= d
set a Difficult!		= <nil>
set h.a list boolean	= <nil>
set h.f list float		= <nil>
set h.b byte size 10T	= <nil>
get a.b.c def:example	= example
get a Difficult!		= Difficult!
get a list nil			= []
get h.a list boolean	= [false true false]
get h.f list float		= [1.2 2.3 3.4]
get h float not exist	= 0
set h.b byte size 10t	= 10995116277760
get b.d def:example		= example
get b.d []object 1->4	= [1 2 3 4]
get b.d []string nil	= []
get b.d []int 1->4		= [1 2 3 4]
set b.d ["1","2","3"]	= <nil>
get b.d []string 1->3	= [1 2 3]
last dump	=
{"a":"Difficult!","b":{"c":{"d":"d","e":"Correct","f":2,"g":false,"t":"1day"},"d":["1","2","3"]},"h":{"a":[false,true,false],"b":"10T","f":[1.2,2.3,3.4]}}
```


### Run yml.go

```bash
&{example.yml 3 0x22d818 {{0 0} 0 0 0 0} map[a:Easy! b:map[c:map[e:Just Do it f:2 g:true t:1day] d:[3 4]] h:1.01]}
get a Easy!				= Easy!
get a.b.c def:example	= example
get b.c.e				= Just Do it
get h is 1.01			= 1.01
get b.c.f def:100,but 2	= 2
get b.c.e not exist		= 0
get b.c.g ON,return T	= true
get b.c.x def:true		= true
get b config			= &{ 0 0x22d818 {{0 0} 0 0 0 0} map[b:map[c:map[e:Just Do it f:2 g:true t:1day] d:[3 4]]]}
get b.d list 3->4		= [3 4]
get b.c.t time:1day		= 24h0m0s
set a.b.c Correct		= <nil>
set b.c.e Correct		= <nil>
set b.c.d d				= <nil>
set b.c.g false 		= <nil>
set b.d list 1->4		= <nil>
get a def:example		= example
get a interface			= map[b:map[c:Correct]]
get a.b.c set Correct	= Correct
get b.c.e set Correct	= Correct
get b.c.g set false		= false
get b.c.d set d			= d
set a Difficult!		= <nil>
set h.a list boolean	= <nil>
set h.f list float		= <nil>
set h.b byte size 10T	= <nil>
get a.b.c def:example	= example
get a Difficult!		= Difficult!
get a list nil			= []
get h.a list boolean	= [false true false]
get h.f list float		= [1.2 2.3 3.4]
get h float not exist	= 0
set h.b byte size 10t	= 10995116277760
get b.d def:example		= example
get b.d []object 1->4	= [1 2 3 4]
get b.d []string nil	= []
get b.d []int 1->4		= [1 2 3 4]
set b.d ["1","2","3"]	= <nil>
get b.d []string 1->3	= [1 2 3]
last dump	=
a: Difficult!
b:
  c:
    d: d
    e: Correct
    f: 2
    g: false
    t: 1day
  d:
  - "1"
  - "2"
  - "3"
h:
  a:
  - false
  - true
  - false
  b: 10T
  f:
  - 1.2
  - 2.3
  - 3.4
```