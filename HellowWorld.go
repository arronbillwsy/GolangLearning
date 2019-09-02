package main

func main()  {
	//buf:= [5]byte{'a','b','c','e','f'}
	//for i:=0;i<5;i++{
	//	//print(buf[i])
	//	//fmt.Print(buf[i])
	//	fmt.Printf("%c", buf[i])
	//	fmt.Println(" ")
	//}
	//fmt.Print(unsafe.Sizeof(buf))
	//fmt.Println(" bytes")
	//host, _ := os.Hostname()
	//addrs, _ := net.LookupIP(host)
	//for _, addr := range addrs {
	//	if ipv4 := addr.To4(); ipv4 != nil {
	//		fmt.Println("IPv4: ", ipv4)
	//		fmt.Println(int(ipv4[3]))
	//	}
	//}
	//fmt.Printf("Hello World")
	//fmt.Println("1+1=", 1+1)
	//fmt.Println(true && false)
	//fmt.Println(true || false)
	//fmt.Println(!true)
	//
	//for i:=1;i<=3;i++ {
	//	fmt.Println(i)
	//}
	//for {
	//	fmt.Println("ok")
	//	break
	//}
	//switch time.Now().Weekday() {
	//case time.Monday,time.Tuesday,time.Wednesday,time.Thursday,time.Friday:
	//	fmt.Println("workday")
	//case time.Saturday,time.Sunday:
	//	fmt.Println("weekday")
	//default:
	//	fmt.Println("happy")
	//}
	//var a [5]int
	//fmt.Println(a)
	//a[4] = 100
	//fmt.Println(a)
	//fmt.Println(a[4])
	//b := [2]int{0,1}
	//fmt.Println(b)
	//var arr [2][3]int
	//for i:=0;i<2;i++{
	//	for j:=0;j<3;j++{
	//		arr[i][j] = i+j
	//	}
	//}
	//fmt.Println(arr)
	//var t [3]string
	//t[1] = "bb"
	//fmt.Println(t)
	//s := make([]string, 3)
	//s[0] = "aa"
	//s[1] = "bb"
	//s[2] = "cc"
	//fmt.Println(s)
	//c:=make([]string, len(s))
	//copy(c, s)
	//fmt.Println(c)
	//s = append(s, "dd")
	//fmt.Println(s)
	//arr := make([][]int, 3)
	//for i:=0;i<len(arr);i++{
	//	arr[i] = make([]int, 3)
	//	for j:=0;j<len(arr[i]);j++{
	//		arr[i][j] = i+j
	//	}
	//}
	//fmt.Println(arr)
	//mymap := make(map[string]int)
	//mymap["key1"] = 1
	//mymap["key2"] = 2
	//mymap["key3"] = 3
	//mymap["key4"] = 4
	//mymap["key5"] = 5
	//mymap["key2"] = 20
	//value1 := mymap["key1"]
	//fmt.Println(value1)
	//fmt.Println(mymap)
	//_, check:=mymap["key1"]
	//fmt.Println(check)
	//delete(mymap, "key1")
	//_, check =mymap["key1"]
	//fmt.Println(check)
	//fmt.Println(mymap)
	//nums:=[]int{1,2,3,4}
	//sum:=0
	//for _,num:=range nums{
	//	sum+=num
	//}
	//fmt.Println(sum)
	//for k,v:=range mymap{
	//	fmt.Println(k,v)
	//}
	//for i,c:=range "this is a test string"{
	//	fmt.Print(i)
	//	fmt.Printf("%c", c)
	//	fmt.Println()
	//}
	//fmt.Println(add(1,2))
	//fmt.Println(multiadd(1,2,3))
	//fmt.Println(cal(3,2,1))
	//fmt.Println(sum(1,2))
	//fmt.Println(sum(1,2,3,4,5))
	//nums:=[]int{1,2,3,4}
	//fmt.Println(sum(nums...))
	//nextInt:=intSeq()
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//nextInts:=intSeq()
	//fmt.Println(nextInts())
	//fmt.Println(fact(10))
	//i:=1
	//fmt.Println("initial", i)
	//zeroval(i)
	//fmt.Println("zeroval", i)
	//zeroptr(&i)
	//fmt.Println("zeroptr", i)
	//fmt.Println("pointer", &i)
	//s:=person{
	//	name: "wsy",
	//	age:  26,
	//}
	//fmt.Println(s.name)
	//sp:=&s
	//fmt.Println(sp.name)
	//sp.age = 18
	//fmt.Println(sp.age)
	//s.age = 10
	//fmt.Println(s.age)
	//r:=rect{
	//	width:  10,
	//	height: 5,
	//}
	//fmt.Println(r.area())
	//fmt.Println(r.perim())
	//rp:=&r
	//fmt.Println(rp.area())
	//fmt.Println(rp.perim())
	//r:=rect{
	//	width:  3,
	//	height: 4,
	//}
	//c:=circle{radius:5}
	//fmt.Println(measure(r))
	//fmt.Println(measure(c))
	//for _,i:=range []int{7,42}{
	//	if r,e:=f1(i);e!=nil{
	//		fmt.Println("f1 failed", e)
	//	}else {
	//		fmt.Println("f1 worked", r)
	//	}
	//	if r,e:=f2(i);e!=nil{
	//		fmt.Println("f2 failed", e)
	//	}else {
	//		fmt.Println("f2 worked", r)
	//	}
	//}
	//_, e:= f2(42)
	//if ae, ok:=e.(*argError);ok{
	//	fmt.Println(ae.arg)
	//	fmt.Println(ae.prob)
	//}
	//f("direct")
	//go f("goroutine")
	//go func(msg string){
	//	fmt.Println(msg)
	//}("going")
	//fmt.Scanln()
	//fmt.Println("done")
	//message:=make(chan string)
	//go func() {message<-"ping"}()
	//msg:=<-message
	//fmt.Println(msg)
	//messages:=make(chan string,2)
	//messages<-"buffered"
	//messages<-"channel"
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)
	//done:=make(chan bool,1)
	//go worker(done)
	//<-done
	//pings := make(chan string,1)
	//pongs := make(chan string,1)
	//ping(pings,"passed message")
	//pong(pings, pongs)
	//fmt.Println(<-pongs)
	//c1:=make(chan string)
	//c2:=make(chan string)
	//go func() {
	//	time.Sleep(time.Second)
	//	c1<-"one"
	//}()
	//go func() {
	//	time.Sleep(time.Second*2)
	//	c2<-"two"
	//}()
	//for i:=0;i<2;i++{
	//	select {
	//	case msg1:=<-c1:
	//		fmt.Println("ok", msg1)
	//	case msg2:=<-c2:
	//		fmt.Println("ok", msg2)
	//	}
	//}
	//jobs:=make(chan int, 5)
	//done:=make(chan bool)
	//go func() {
	//	j, more:=<-jobs
	//	if more{
	//		fmt.Println("received job", j)
	//	}else{
	//		fmt.Println("received all jobs")
	//		done<-true
	//		return
	//	}
	//}()
	//for j:=1;j<=3;j++{
	//	jobs<-j
	//	fmt.Println("sent job", j)
	//}
	//close(jobs)
	//fmt.Println("sent all jobs")
	//<-done
	//queue:=make(chan string,2)
	//queue<-"one"
	//queue "two"
	//close(queue)
	//for element:=range queue{
	//	fmt.Println()
	//}
	//jobs:=make(chan int, 100)
	//results:=make(chan int, 100)
	//for w:=1;w<=3;w++{
	//	go worker(w, jobs, results)
	//}
	//for j:=1;j<=5;j++{
	//	jobs<-j
	//}
	//for a:=1;a<=5;a++{
	//	<-results
	//}
	//var wg sync.WaitGroup
	//for i:=1;i<=5;i++{
	//	wg.Add(1)
	//	go worker(i,&wg)
	//}
	//wg.Wait()
	//requests:=make(chan int, 5)
	//for i:=1;i<=5;i++{
	//	requests<-i
	//}
	//close(requests)
	//limiter:=time.Tick(200*time.Millisecond)
	//for req:=range requests{
	//	<-limiter
	//	fmt.Println("request", req, time.Now())
	//}
	//burstyLimiter:=make(chan time.Time,3)
	//for i:=0;i<3;i++{
	//	burstyLimiter<- time.Now()
	//}
	//go func() {
	//	for t:=range time.Tick(200*time.Millisecond){
	//		burstyLimiter<-t
	//	}
	//}()
	//burstyRequests:=make(chan int, 5)
	//for i:=1;i<=5;i++{
	//	burstyRequests<-i
	//}
	//close(burstyRequests)
	//for req:=range burstyRequests{
	//	<-burstyLimiter
	//	fmt.Println("request", req, time.Now())
	//}
	//var ops uint64
	//for i:=0;i<50;i++{
	//	go func() {
	//		for {
	//			atomic.AddUint64(&ops,1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//time.Sleep(time.Second)
	//opsFinal:=atomic.LoadUint64(&ops)
	//fmt.Println("ops", opsFinal)
	//var state = make(map[int]int)
	//var mutex = &sync.Mutex{}
	//var readOps uint64
	//var writeOps uint64
	//for r:=0;r<100;r++{
	//	go func() {
	//		total:=0
	//		for{
	//			key:=rand.Intn(5)
	//			mutex.Lock()
	//			total+=state[key]
	//			mutex.Unlock()
	//			atomic.AddUint64(&readOps,1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//for w:=0;w<10;w++{
	//	go func() {
	//		for {
	//			key:=rand.Intn(5)
	//			val:=rand.Intn(100)
	//			mutex.Lock()
	//			state[key] = val
	//			mutex.Unlock()
	//			atomic.AddUint64(&writeOps,1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//time.Sleep(time.Second)
	//readOpsFinal:=atomic.LoadUint64(&readOps)
	//fmt.Println("readOps", readOpsFinal)
	//writeOpsFinal:=atomic.LoadUint64(&writeOps)
	//fmt.Println("writeOps", writeOpsFinal)
	//mutex.Lock()
	//fmt.Println("state", state)
	//mutex.Unlock()
	//var readOps uint64
	//var writeOps uint64
	//reads:=make(chan readOp)
	//writes:=make(chan writeOp)
	//go func() {
	//	var state = make(map[int]int)
	//	for {
	//		select {
	//		case read:=<-reads:
	//			read.resp<-state[read.key]
	//		case write := <-writes:
	//			state[write.key] = write.val
	//			write.resp<-true
	//		}
	//	}
	//}()
	//for r:=0;r<100;r++{
	//	go func() {
	//		for {
	//			read:=readOp{
	//				key:  rand.Intn(5),
	//				resp: make(chan int),
	//			}
	//			reads<-read
	//			<-read.resp
	//			atomic.AddUint64(&readOps,1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//for w:=0;w<10;w++{
	//	go func() {
	//		for {
	//			write:=writeOp{
	//				key:  rand.Intn(5),
	//				val:  rand.Intn(100),
	//				resp: make(chan bool),
	//			}
	//			writes<-write
	//			<-write.resp
	//			atomic.AddUint64(&writeOps,1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//time.Sleep(time.Second)
	//readOpsFinal:=atomic.LoadUint64(&readOps)
	//fmt.Println("readOps", readOpsFinal)
	//writeOpsFinal:=atomic.LoadUint64(&writeOps)
	//fmt.Println("writeOps", writeOpsFinal)
	//
	//strs:=[]string{"c","a","b"}
	//sort.Strings(strs)
	//fmt.Println("strings:",strs)
	//ints:=[]int{7,2,4}
	//sort.Ints(ints)
	//fmt.Println("ints:",ints)
	//s:=sort.IntsAreSorted(ints)
	//fmt.Println("sorted",s)
	//fruits:=[]string{"peach","banana","kiwi"}
	//sort.Sort(byLength(fruits))
	//fmt.Println(fruits)
	//panic("a problem")
	//_,err:=os.Create("/tmp/file")
	//if err!=nil{
	//	panic(err)
	//}
	//f:=createFile("D:/Users/siyuan_wu/go/src/GOLearning/resources/defer.txt")
	//defer closeFile(f)
	//writeFile(f)
	//var strs = []string{"peach", "apple", "pear", "plum"}
	//fmt.Println(Index(strs, "pear"))
	//fmt.Println(Include(strs,"grape"))
	//fmt.Println(Any(strs, func(v string) bool {
	//	return strings.HasPrefix(v,"p")
	//}))
	//fmt.Println(All(strs, func(v string) bool {
	//	return strings.HasPrefix(v,"p")
	//}))
	//fmt.Println(Filter(strs, func(v string) bool {
	//	return strings.Contains(v,"e")
	//}))
	//fmt.Println(Map(strs, strings.ToUpper))
	//bolB,_:=json.Marshal(true)
	//fmt.Println(string(bolB))
	//slcD:= []string{"peach", "apple", "pear", "plum"}
	//slcb,_:=json.Marshal(slcD)
	//fmt.Println(string(slcb))
	//s:="postgres://user:pass@host.com:5432/path?k=v#f"
	//u, err:=url.Parse(s)
	//if err!=nil{
	//	panic(err)
	//}
	//fmt.Println(u.Scheme)
	//fmt.Println(u.User.Username())
	//p,_:=u.User.Password()
	//fmt.Println(p)
	//fmt.Println(u.Host)
	//host, port,_:=net.SplitHostPort(u.Host)
	//fmt.Println(host)
	//fmt.Println(port)
	//fmt.Println(u.Path)
	//fmt.Println(u.Fragment)
	//fmt.Println(u.RawQuery)
	//m,_:=url.ParseQuery(u.RawQuery)
	//fmt.Println(m)
	//fmt.Println(m["k"][0])
	//s:="sha1 this string"
	//h:=sha1.New()
	//h.Write([]byte(s))
	//bs:=h.Sum(nil)
	//fmt.Println(s)
	//fmt.Printf("%x\n", bs)
	//
	//dat, err:=ioutil.ReadFile("D:/Users/siyuan_wu/go/src/GOLearning/resources/defer.txt")
	//check(err)
	//fmt.Println(string(dat))
	//f, err:=os.Open("D:/Users/siyuan_wu/go/src/GOLearning/resources/defer.txt")
	//check(err)
	//b1:=make([]byte,5)
	//n1, err:=f.Read(b1)
	//check(err)
	//fmt.Printf("%d bytes:%s\n", n1,string(b1[:n1]))
	//d1:=[]byte("hello world\n")
	//err:=ioutil.WriteFile("D:/Users/siyuan_wu/go/src/GOLearning/resources/defer.txt", d1, 0644)
	//check(err)
	//scanner:=bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	ucl:=strings.ToUpper(scanner.Text())
	//	fmt.Println(ucl)
	//}
	//if err:=scanner.Err();err!=nil {
	//	fmt.Fprintf(os.Stderr, "error", err)
	//	os.Exit(1)
	//}
	//argsWithProg:=os.Args
	//fmt.Println(argsWithProg)
	//resp, err:=http.Get("http://gobyexample.com")
	//if err!=nil {
	//	panic(err)
	//}
	//fmt.Println("response status:", resp.Status)
	//scanner:=bufio.NewScanner(resp.Body)
	//for i:=0;scanner.Scan()&&i<5;i++{
	//	fmt.Println(scanner.Text())
	//}
	//if err:=scanner.Err();err!=nil {
	//	panic(err)
	//}
	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/headers", headers)
	//http.ListenAndServe(":8090", nil)

}
//
//
//
//func hello(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "hello\n")
//}
//
//func headers(w http.ResponseWriter, req *http.Request) {
//	for name,header:=range req.Header {
//		for _,h:=range header{
//			fmt.Fprintf(w, "%v:%v\n", name, h)
//		}
//	}
//}
//
//func check(e error) {
//	if e!=nil {
//		panic(e)
//	}
//}
//
//type response1 struct {
//	Page int
//	Fruits []string
//}
//type response2 struct {
//	Page int `json:"page"`
//	Fruits []string `json:"fruits"`
//}
//
//func Index(vs []string, t string) int {
//	for i, v := range vs {
//		if v == t {
//			return i
//		}
//	}
//	return -1
//}
//
//func Include(vs []string, t string) bool {
//	return Index(vs, t) >= 0
//}
//
//func Any(vs []string, f func(string) bool) bool {
//	for _, v := range vs {
//		if f(v) {
//			return true
//		}
//	}
//	return false
//}
//
//func All(vs []string, f func(string) bool) bool {
//	for _, v := range vs {
//		if !f(v) {
//			return false
//		}
//	}
//	return true
//}
//
//func Filter(vs []string, f func(string) bool) []string {
//	vsf := make([]string, 0)
//	for _, v := range vs {
//		if f(v) {
//			vsf = append(vsf, v)
//		}
//	}
//	return vsf
//}
//
//func Map(vs []string, f func(string) string) []string {
//	vsm := make([]string, len(vs))
//	for i, v := range vs {
//		vsm[i] = f(v)
//	}
//	return vsm
//}
//
//func createFile(p string) *os.File {
//	fmt.Println("creating")
//	f, err:=os.Create(p)
//	if err!=nil{
//		panic(err)
//	}
//	return f
//}
//
//func writeFile(f *os.File)  {
//	fmt.Println("writing")
//	fmt.Fprintln(f,"data")
//}
//
//func closeFile(f *os.File)  {
//	fmt.Println("closing")
//	err:=f.Close()
//	if err!=nil{
//		fmt.Fprintf(os.Stderr, "error: %v\n", err)
//		os.Exit(1)
//	}
//}
//
//type byLength []string
//
//func (s byLength)Len() int {
//	return len(s)
//}
//
//func (s byLength)Swap(i,j int)  {
//	s[i],s[j] = s[j],s[i]
//}
//
//func (s byLength)Less(i,j int) bool {
//	return len(s[i])<len(s[j])
//}
//
//type readOp struct {
//	key int
//	resp chan int
//}
//
//type writeOp struct {
//	key int
//	val int
//	resp chan bool
//}
//
//func worker(id int, wg *sync.WaitGroup){
//	fmt.Printf("worker %d starting\n", id)
//	time.Sleep(time.Second)
//	fmt.Println("worder %d done", id)
//	wg.Done()
//}
//
//func ping(pings chan<- string, msg string)  {
//	pings<-msg
//}
//
//func pong(pings <-chan string, pongs chan <- string)  {
//	msg:= <- pings
//	pongs<-msg
//}
//
////func worker(id int, jobs <-chan int, results chan<- int) {
////	for j:=range jobs{
////		fmt.Println("wokder", id, "start",j)
////		time.Sleep(time.Second)
////		fmt.Println("worker", id,"finished job",j)
////		results<-j*2
////	}
////}
//
////func worker(done chan bool)  {
////	fmt.Print("working")
////	time.Sleep(time.Second)
////	fmt.Println("done")
////	done<-true
////}
//
//func f(from string)  {
//	for i:=0; i<3; i++{
//		fmt.Println(from, ":", i)
//	}
//}
//
//func f1(arg int) (int, error) {
//	if arg == 42{
//		return -1, errors.New("cannot work with 42")
//	}
//	return arg+3, nil
//}
//
//type argError struct {
//	arg int
//	prob string
//}
//
//func (e *argError) Error() string {
//	return fmt.Sprintf("%d - %s", e.arg, e.prob)
//}
//
//func f2(arg int) (int, error) {
//	if arg == 42{
//		return -1, &argError{
//			arg:  arg,
//			prob: "cannot work with it",
//		}
//	}
//	return arg+3, nil
//}
//
//type geometry interface {
//	area() float64
//	perim() float64
//}
//
//func measure(g geometry) (area float64, perim float64) {
//	//fmt.Println(g)
//	//fmt.Println(g.area())
//	//fmt.Println(g.perim())
//	return g.area(), g.perim()
//}
//
//type circle struct {
//	radius float64
//}
//
//type rect struct {
//	width, height float64
//}
//
////func (r *rect) area() int  {
////	return r.width*r.height
////}
//
//func (r rect) perim() float64 {
//	return 2*r.width+2*r.height
//}
//
//func (r rect) area() float64 {
//	return r.width*r.height
//}
//
//func (c circle) area() float64 {
//	return math.Pi*c.radius*c.radius
//}
//
//func (c circle) perim() float64 {
//	return 2*math.Pi*c.radius
//}
//
//type person struct {
//	name string
//	age int
//}
//
//func zeroval(ival int)  {
//	ival = 0
//}
//
//func zeroptr(iptr *int)  {
//	*iptr = 0
//}
//
//func fact(n int) int {
//	if n ==0 {
//		return 1
//	}
//	return n*fact(n-1)
//}
//
//func intSeq() func()int {
//	i:=0
//	return func() int {
//		i++
//		return i
//	}
//}
//
//func add(a int, b int) int {
//	return a+b
//}
//
//func multiadd(a,b,c int) int {
//	return a+b+c
//}
//
//func cal(a,b,c int) (int, int){
//	return a-b-c, a+b+c
//}
//
//func sum(nums ...int) int {
//	fmt.Println(nums, " ")
//	total:=0
//	for _,num:=range nums{
//		total+=num
//	}
//	return total
//}
