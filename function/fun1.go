package function

import (
	"learn/common"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 160. Intersection of Two Linked Lists
func GetIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	dict := make(map[*common.ListNode]struct{})
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := dict[headA]; ok {
				return headA
			} else {
				dict[headA] = struct{}{}
			}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := dict[headB]; ok {
				return headB
			} else {
				dict[headB] = struct{}{}
			}
			headB = headB.Next
		}
	}
	return nil
}

// 162. Find Peak Element
func FindPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}

	res := 1
	for res < len(nums) {
		rcode := 0
		if res == len(nums)-1 {
			rcode = 1
		} else if nums[res] > nums[res+1] {
			rcode = 2
		} else {
			rcode = 3
		}

		if nums[res] > nums[res-1] && (rcode == 1 || rcode == 2) {
			break
		}
		if rcode == 2 {
			res += 2
		} else {
			res++
		}
	}
	return res
}

// 164. Maximum Gap
func MaximumGap(nums []int) int {
	// 3,6,9,1 -> 3
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

// 165. Compare Version Numbers
func CompareVersion(version1 string, version2 string) int {
	// 1.0.0.0001
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	i := 0
	for ; i < len(v1) && i < len(v2); i++ {
		vv1, _ := strconv.Atoi(v1[i])
		vv2, _ := strconv.Atoi(v2[i])
		if vv1 > vv2 {
			return 1
		}
		if vv1 < vv2 {
			return -1
		}
	}

	v3 := v1
	res := 1
	if i == len(v1) {
		v3 = v2
		res = -1
	}

	for ; i < len(v3); i++ {
		vv3, _ := strconv.Atoi(v3[i])
		if vv3 != 0 {
			return res
		}
	}

	return 0
}

// 167. Two Sum II - Input array is sorted
func TwoSum2(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

// 166. Fraction to Recurring Decimal
func FractionToDecimal(numerator int, denominator int) string {
	// assert denominator != 0
	res := ""
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		numerator *= -1
		res += "-"
	}
	red := (numerator % denominator) * 10
	res += strconv.Itoa(numerator / denominator)
	if red != 0 {
		res += "."
	}
	numMap := make(map[int]int)
	index := len(res)

	for red != 0 {
		numMap[red] = index
		quo := red / denominator
		index++
		res += strconv.Itoa(quo)
		red %= denominator
		red *= 10

		if v, ok := numMap[red]; ok {
			res = res[:v] + "(" + res[v:]
			res += ")"
			break
		}
	}

	return res
}

// 168. Excel Sheet Column Title
func ConvertToTitle(n int) string {
	res := ""
	for n > 0 {
		if n == common.BIG_ENGLISH_CHAR_NUM {
			res = "Z" + res
			return res
		}

		dev := n % common.BIG_ENGLISH_CHAR_NUM
		if dev > 0 {
			res = string('A'+dev-1) + res
		} else {
			res = "Z" + res
			n -= common.BIG_ENGLISH_CHAR_NUM
		}
		n /= common.BIG_ENGLISH_CHAR_NUM
	}
	return res
}

// 169. Majority Element
func MajorityElement(nums []int) int {
	lastNum := nums[0]
	lastAce := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != lastNum {
			lastAce--
			if lastAce == 0 {
				lastNum = nums[i]
				lastAce = 1
			}
		} else {
			lastAce++
		}
	}
	return lastNum
}

// 171. Excel Sheet Column Number
func TitleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= common.BIG_ENGLISH_CHAR_NUM
		res += int(s[i]-'A') + 1
	}
	return res
}

// 172. Factorial Trailing Zeroes
func TrailingZeroes(n int) int {
	sub := 5
	res := 0
	for n >= sub {
		res += n / sub
		sub *= 5
	}
	return res
}

// 174. Dungeon Game
func CalculateMinimumHP(dungeon [][]int) int {
	type hpHistory struct {
		hp  int
		min int
	}

	n := len(dungeon)
	m := len(dungeon[0])
	dpDungeon := make([][][]hpHistory, n)
	for i := 0; i < n; i++ {
		dpDungeon[i] = make([][]hpHistory, m)
	}
	dpDungeon[0][0] = append(dpDungeon[0][0], hpHistory{hp: dungeon[0][0], min: dungeon[0][0]})

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			curDp := dpDungeon[i][j]
			if i > 0 {
				for z := 0; z < len(dpDungeon[i-1][j]); z++ {
					curDp = append(curDp, hpHistory{hp: dpDungeon[i-1][j][z].hp + dungeon[i][j]})
					curDp[len(curDp)-1].min = int(math.Min(float64(curDp[len(curDp)-1].hp), float64(dpDungeon[i-1][j][z].min)))
				}
			}
			if j > 0 {
				for z := 0; z < len(dpDungeon[i][j-1]); z++ {
					isAdd := true
					newDp := hpHistory{
						hp:  dpDungeon[i][j-1][z].hp + dungeon[i][j],
						min: int(math.Min(float64(dpDungeon[i][j-1][z].hp+dungeon[i][j]), float64(dpDungeon[i][j-1][z].min))),
					}

					for s := 0; s < len(curDp); s++ {
						if newDp.hp <= curDp[s].hp && newDp.min <= curDp[s].min {
							isAdd = false
							break
						}
						if newDp.hp >= curDp[s].hp && newDp.min >= curDp[s].min {
							curDp[s].hp = newDp.hp
							curDp[s].min = newDp.min
							isAdd = false
							break
						}
					}
					if isAdd {
						curDp = append(curDp, newDp)
					}
				}
			}
			dpDungeon[i][j] = curDp
		}
	}
	res := common.MAXINTNUM
	for i := 0; i < len(dpDungeon[n-1][m-1]); i++ {
		if dpDungeon[n-1][m-1][i].min > 0 {
			return 1
		}
		if -dpDungeon[n-1][m-1][i].min+1 < res {
			res = -dpDungeon[n-1][m-1][i].min + 1
		}
	}
	return res
}

// 175. Combine Two Tables # cool: sql question ?
// select FirstName, LastName, City, State from Person left join Address using (PersonId)

// 176. Second Highest Salary
// select IFNULL((select distinct Salary as SecondHighestSalary from Employee order by Salary desc limit 1,1), NULL) as SecondHighestSalary

// 177. Nth Highest Salary
/*
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N = N-1;
  RETURN (
	  select IFNULL((select distinct Salary as SecondHighestSalary from Employee order by Salary desc limit  N,1), NULL)
  );
END
*/

// 178. Rank Scores
/*
SELECT Score, rr as 'Rank' from
(SELECT s1.Score, count(distinct(s2.Score)) as rr
FROM Scores s1,Scores s2
WHERE s1.Score<=s2.Score
GROUP BY s1.Id) as t1
ORDER BY rr
*/

// 179. Largest Number
type dumpNumber struct {
	childs map[int]*dumpNumber
	vals   [10]string
}

func LargestNumber(nums []int) string {
	sortDump := &dumpNumber{childs: make(map[int]*dumpNumber)}
	for i := 0; i < len(nums); i++ {
		curDump := sortDump
		curNum := strconv.Itoa(nums[i])
		for j := 0; j < 10; j++ {
			index := int(curNum[j%len(curNum)] - '0')
			if _, ok := curDump.childs[index]; !ok {
				curDump.childs[index] = &dumpNumber{childs: make(map[int]*dumpNumber)}
			}
			curDump = curDump.childs[index]
		}
		curDump.vals[len(curNum)-1] += curNum
	}

	i := 0
	res := backDump(sortDump, 0)
	for i = 0; i < len(res); i++ {
		if res[i] != '0' {
			break
		}
	}
	if i == len(res) {
		return "0"
	}
	return res[i:]
}

func backDump(dump *dumpNumber, n int) string {
	res := ""
	childs := dump.childs
	if n == 10 {
		// 叶子层
		for i := 0; i <= 9; i++ {
			res = res + dump.vals[i]
		}
	} else {
		for i := 9; i >= 0; i-- {
			if child, ok := childs[i]; ok {
				n++
				res = res + backDump(child, n)
				n--
			}
		}
	}
	return res
}

// 180. Consecutive Numbers
/*
SELECT distinct Num as ConsecutiveNums
from (
  select Num,
    case
      when @prev = Num then @count := @count + 1
      when (@prev := Num) is not null then @count := 1
    end as CNT
  from Logs, (select @prev := null,@count := null) as t
) as temp
where temp.CNT >= 3
*/

// 181. Employees Earning More Than Their Managers
/*
SELECT a.Name as 'Employee'
FROM Employee a,Employee b
WHERE a.ManagerId IS NOT NULL and b.Id = a.ManagerId and a.Salary > b.Salary
*/

// 182. Duplicate Emails
/*
SELECT DISTINCT(a.Email) as 'Email'
FROM Person a, Person b
WHERE a.Email = b.Email and a.Id != B.Id
*/

// 183. Customers Who Never Order
/*
SELECT Name as 'Customers'
FROM Customers
WHERE Id
not in (
    SELECT DISTINCT(CustomerId) FROM Orders
)
*/

// 184. Department Highest Salary
/*
SELECT Department.Name as Department, Employee.Name as Employee, Employee.Salary as Salary
FROM Department LEFT JOIN Employee
ON Department.Id = Employee.DepartmentId
WHERE (Employee.Salary, Employee.DepartmentId) IN (SELECT Max(Salary), DepartmentId From Employee GROUP BY DepartmentId)
*/

// 185. Department Top Three Salaries
/*
SELECT T2.Name AS "Department", T1.Name AS "Employee", T1.Salary AS "Salary"
FROM
(
    SELECT *
    FROM Employee
    WHERE (DepartmentId, Salary) IN (
        SELECT a.DepartmentId, a.Salary FROM (
            SELECT * FROM Employee GROUP BY Salary, DepartmentId
        ) as a
        WHERE (
            SELECT count(*) FROM
            (
                SELECT * FROM Employee GROUP BY Salary, DepartmentId
            ) AS b WHERE a.DepartmentId = b.DepartmentId AND a.Salary < b.Salary
        ) < 3
    )
) AS T1 LEFT JOIN Department AS T2 ON T1.DepartmentId = T2.Id
WHERE T2.Name IS NOT NULL
*/

// 187. Repeated DNA Sequences
func FindRepeatedDnaSequences(s string) []string {
	transNum := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	transByte := []byte{'A', 'C', 'G', 'T'}
	transStr := func(v int) string {
		res := ""
		for len(res) < 10 {
			res = string(transByte[v%4]) + res
			v /= 4
		}
		return res
	}
	DNAMap := make(map[int]bool)
	curSize := 0
	curDNA := 0
	for i := 0; i < len(s); i++ {
		if curSize < 10 {
			curDNA = (curDNA << 2) + transNum[s[i]]
			curSize++
		} else {
			curDNA = curDNA % (1 << 18)
			curDNA = (curDNA << 2) + transNum[s[i]]
		}
		if curSize == 10 {
			if _, ok := DNAMap[curDNA]; ok {
				DNAMap[curDNA] = true
			} else {
				DNAMap[curDNA] = false
			}
		}
	}

	res := []string{}
	for DNA, v := range DNAMap {
		if v {
			res = append(res, transStr(DNA))
		}
	}
	return res
}

// 188. Best Time to Buy and Sell Stock IV
func MaxProfitV4(k int, prices []int) int {
	if k == 0 {
		return 0
	}

	if k >= len(prices)/2 {
		// Code is deleted by myself
		// return MaxProfitV2(prices)
	}

	pricesMatrix := make([][2]int, k)
	for i := 0; i < k; i++ {
		pricesMatrix[i][0] = common.MININTNUM
	}

	for i := 0; i < len(prices); i++ {
		pricesMatrix[0][0] = int(math.Max(float64(pricesMatrix[0][0]), float64(-prices[i])))
		pricesMatrix[0][1] = int(math.Max(float64(pricesMatrix[0][1]), float64(pricesMatrix[0][0]+prices[i])))
		for j := 1; j < k; j++ {
			pricesMatrix[j][0] = int(math.Max(float64(pricesMatrix[j][0]), float64(pricesMatrix[j-1][1]-prices[i])))
			pricesMatrix[j][1] = int(math.Max(float64(pricesMatrix[j][1]), float64(pricesMatrix[j][0]+prices[i])))
		}
	}

	return pricesMatrix[k-1][1]
}

// 189. Rotate Array
func RotateArr(nums []int, k int) {
	k = k % len(nums)
	a := make([]int, k)
	for i := len(nums) - 1; i >= 0; i-- {
		if len(nums)-i <= k {
			a[k-len(nums)+i] = nums[i]
		}

		if i >= k {
			nums[i] = nums[i-k]
		} else {
			nums[i] = a[i]
		}
	}
}

// 190. Reverse Bits
func ReverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		res += num & 1
		num >>= 1
	}
	return res
}

// 191. Number of 1 Bits
func HammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

// 192. Word Frequency
/*
declare -A m_dict
while read line;

do
	# 切割字符传
	array=(${line// / })
	for var in ${array[@]}
	if [ ! -n m_dict[${var}] ]; then
		m_dict[${var}] = 1
	else
		m_dict[${var}] = $m_dict[${var}] + 1

done < word.txt

for key in $(echo ${!m_dict[*]})
do
    echo "$key ${dic[$key]}"
done
*/

// 193. Valid Phone Numbers
/*
grep -P '^(\d{3}-|\(\d{3}\) )\d{3}-\d{4}$' file.txt
*/

// 194.Transpose File
/*
transpose=`head -n1 file.txt | wc -w`

for i in `seq 1 $transpose`
do
    echo `cut -d' ' -f$i file.txt`
done
*/

// 195. Tenth Line
/*
sed -n '10p' file.txt
*/

// 196. Delete Duplicate Emails
/*
DELETE FROM `Person` WHERE `Id` NOT IN (
    SELECT * FROM (SELECT Min(`Id`) FROM `Person` GROUP BY (`Email`)) As b
)
*/

// 197. Rising Temperature
/*
select a.Id from  Weather as a join Weather as b on a.Temperature > b.Temperature and dateDiff(a.RecordDate,b.RecordDate) = 1
*/

// 198. House Robber
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	UndoNums := make([]int, len(nums))
	DoNums := make([]int, len(nums))
	DoNums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		UndoNums[i] = DoNums[i-1]
		if UndoNums[i-1]+nums[i] > DoNums[i-1] {
			DoNums[i] = UndoNums[i-1] + nums[i]
		} else {
			DoNums[i] = DoNums[i-1]
		}
	}

	return DoNums[len(nums)-1]
}

// 199. Binary Tree Right Side View
func RightSideView(root *common.TreeNode) []int {
	res := []int{}
	rightSideViewSub(root, 0, &res)
	return res
}

func rightSideViewSub(root *common.TreeNode, high int, res *[]int) {
	if root == nil {
		return
	}

	if high >= len(*res) {
		*res = append(*res, root.Val)
	}
	if root.Right != nil {
		rightSideViewSub(root.Right, high+1, res)
	}
	if root.Left != nil {
		rightSideViewSub(root.Left, high+1, res)
	}
	return
}

// 200. Number of Islands
func NumIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if numIslandsSub(grid, i, j) {
				res++
			}
		}
	}
	return res
}

func numIslandsSub(grid [][]byte, i, j int) bool {
	if grid[i][j] == 1 {
		grid[i][j] = 0
		if i < len(grid)-1 {
			numIslandsSub(grid, i+1, j)
		}
		if i > 0 {
			numIslandsSub(grid, i-1, j)
		}
		if j > 0 {
			numIslandsSub(grid, i, j-1)
		}
		if j < len(grid[0])-1 {
			numIslandsSub(grid, i, j+1)
		}
		return true
	}

	return false
}
