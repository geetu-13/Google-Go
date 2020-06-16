package main
import(
	"fmt"
	"sort"
	"sync"
)
func min(a,b int)(int){
	if (a<b){
		return a
	}
	return b
}
func pos(a int)(int){
	if (a<0){
		a=a*(-1)
	}
	return a
}

func merge(a[]int, b[]int)([]int){
	l:=len(a)+len(b)
	an:=0
	bn:=0
	x:=make([]int, l)
	i:=0
	for (i<l){
		if (an<len(a) && bn<len(b)){
			if (a[an]<b[bn]){
				x[i]=a[an]
				an=an+1
			}else{
				x[i]=b[bn]
				bn=bn+1
			}
		} else if (bn<len(b)){
			x[i]=b[bn]
			bn=bn+1
		} else{
			x[i]=a[an]
			an=an+1
		}
		
		i=i+1
	}
	return x
}

func sorted_array(wg *sync.WaitGroup,sli[]int){
	sort.Ints(sli)
	fmt.Println(sli)
	wg.Done()
}

func part (x int)(int){
	n:=x/4
	a:=n*4
	b:=(n+1)*4
	c:=(n-1)*4
	a=pos(x-a)
	b=pos(x-b)
	c=pos(x-c)
	least:=min(a,b)
	least=min(least,c)
	switch(least){
	case a: return n
	case b: return n+1
	default: return n-1
	}
}
func main(){
	var x int 
	fmt.Println("enter size of array: ")
	fmt.Scan(&x)

	arr:=make([]int, x)
	i:=0
	fmt.Println("enter array elements: ")
	for(i<x){
		fmt.Scan(&arr[i])
		i=i+1
	}	

	fmt.Println("whole array: ")
	fmt.Println(arr)

	part_size:=part(x)
	
	arr1:=make([]int, part_size)
	i=0
	for(i<part_size){
		arr1[i]=arr[i]
		i=i+1
	}

	arr2:=make([]int, part_size)
	j:=0
	for(j<part_size){
		arr2[j]=arr[i]
		i=i+1
		j=j+1
	}

	arr3:=make([]int, part_size)
	j=0
	for(j<part_size){
		arr3[j]=arr[i]
		i=i+1
		j=j+1
	}

	part_size=x-3*part_size
	arr4:=make([]int, part_size)
	j=0
	for(j<part_size){
		arr4[j]=arr[i]
		i=i+1
		j=j+1
	}

	fmt.Println("4 array parts: ")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	fmt.Println("Sorted sub-arrays:")
	var wg sync.WaitGroup
	wg.Add(4)
	go sorted_array(&wg, arr1)
	go sorted_array(&wg, arr2)
	go sorted_array(&wg, arr3)
	go sorted_array(&wg, arr4)

	wg.Wait()
	
	fmt.Println("combined whole array: ")
	a:=merge(arr1,arr2)
	b:=merge(arr3,arr4)
	fmt.Println(merge(a,b))
}