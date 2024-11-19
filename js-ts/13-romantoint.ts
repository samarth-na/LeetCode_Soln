function romanToInt(s: string): number {
	//	I             1
	//V             5
	//X             10
	//L             50
	//C             100
	//D             500
	//M             1000

	let ans: number = 0;
	for (const key of s) {
		switch (key) {
			case "I":
				ans += 1;
				break;
			case "V":
				ans += 5;
				break;
			case "X":
				ans += 10;
				break;
			case "L":
				ans += 50;
				break;
			case "C":
				ans += 100;
				break;
			case "D":
				ans += 500;
				break;
			case "M":
				ans += 1000;
				break;
			default:
				break;
		}
	}

	return ans;
}
let ans = romanToInt("IIII");
console.log(ans);
