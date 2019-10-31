package def

func AssginMent(){

	//x  = 1
	//*p = true
	//person.name = "lbj"
	//count[x] = count[x] * scale

	// v := 1
    // v++
    // v--

    //x , y = y,x
	//a[i],a[j] = a[j],a[i]

   //f, err = os.Open("filepath")

   //v ,ok = m[key]
  // v = m[key]

   //v,ok =x.(T)
   //v = x.(T)

   //v,ok =<-ch
   //v = <-ch

   //_,ok = m[key]
   //_,ok = mm[""],false
   //_ = mm[""]

  //_,err = io.Copy(dst,src)
  //_,ok = x.(T)

    medals  := []string {"a", "b","c"}
    medals[0] = "1"
    medals[1] = "2"
    medals[2] = "3"
}

func gcd(x ,y int) int {

	for y != 0{
		x ,y = y,x%y
	}
	return x
}

func fib(n int) int {

	x,y := 0 ,1
	for i:=0;i<n; i++{
		x,y = y,x+y
	}
	return x
}