func check(nums []int) bool {
    n := len(nums)
    first := 0

    for i := 0; i < n; i++ {
        if nums[(i - 1 + n) % n] > nums[i] {
            first = i
            break
        }
    }

    arr := append(nums, nums...)
    for i := first; i < first + n - 1; i++ {
        if arr[i] > arr[i+1] {
            return false
        }
    }
    return true
}
