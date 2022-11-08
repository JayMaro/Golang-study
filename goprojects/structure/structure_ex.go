package main

import "fmt"

type House struct {
	address   string
	price     int32
	houseType string
}

type User struct {
	name  string
	no    int
	level int
}

type VipUser struct {
	userInfo User
	level    int
}

type VipEmbeddedUser struct {
	User
	level int
}

type MemoryBad struct {
	// 메모리는 64비트 컴퓨터 기준 8바이트 단위로 사용하는 것이 좋기 때문에
	// 메모리 패딩(8바이트 단위로 설정)이 발생하고 따라서 해당 구조체의 크기는
	// 8(1) + 8(1) + 8(1) + 8(8) + 8(8) = 40이 된다.
	A int8
	B int
	C int8
	D int
	E int8
}

type MemoryGood struct {
	// 위와 달리 해당 구조체에서는 앞에 8바이트가 되지 않는 필드들을 하나의 8바이트 안에 몰아넣기 때문에
	// 8(1 + 1 + 1) + 8(8) + 8(8) = 24가 되어 좀 더 메모리를 효율적으로 사용할 수 있다.
	A int8
	B int8
	C int8
	D int
	E int
}

func main() {
	var house1 House = House{"강남구", 100000000, "아파트"}
	house2 := House{address: "강남구", houseType: "아파트"}
	fmt.Println(house1)
	fmt.Println(house2)

	var user1 User = User{"마로", 1, 10}
	fmt.Println("user1 = ", user1)
	var vipUser1 VipUser = VipUser{User{"마시", 2, 11}, 2}
	fmt.Println("vipUser1 = ", vipUser1)
	fmt.Println(
		vipUser1.userInfo.name,
		vipUser1.userInfo.no,
		vipUser1.userInfo.level,
		vipUser1.level,
		vipUser1.userInfo,
	)
	var vipUser2 VipUser = VipUser{user1, 2}
	fmt.Println("vipUser2 = ", vipUser2)
	user1.no = 100
	fmt.Println("diff user1 = ", user1)
	fmt.Println("diff vipUser2 = ", vipUser2) // 값 자체를 복사하기에 user1의 변화가 영향을 주지 않는다.
	var vipEmbeddedUser1 VipEmbeddedUser = VipEmbeddedUser{User{"마시마로", 3, 12}, 3}
	fmt.Println(vipEmbeddedUser1)
	fmt.Println(vipEmbeddedUser1.User)
	fmt.Println(vipEmbeddedUser1.name, vipEmbeddedUser1.level, vipEmbeddedUser1.no, vipEmbeddedUser1.User.level)
}
