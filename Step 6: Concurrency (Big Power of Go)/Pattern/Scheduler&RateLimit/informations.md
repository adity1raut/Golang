Bhai, chalo **Rate Limiter** aur **Scheduler** ke concepts ko simple language mein samajhte hain. Dono concurrency aur system design mein important hain, especially jab resources ko control karna ho ya tasks ko systematically execute karna ho. Ek ek karke dekhte hain:

---

### **Rate Limiter**
**Rate Limiter** ka matlab hai kisi system ke resources ya operations ke usage ko control karna, taaki system overload na ho. Yani, ek specific time frame mein kitne requests ya tasks process honge, uski limit set karna.

#### **Kaise kaam karta hai?**
- Rate Limiter ek rule ya boundary set karta hai, jaise "ek second mein max 100 requests allow karo."
- Agar requests ya tasks is limit se zyada aate hain, toh baki ke requests ko ya toh drop kar diya jata hai, ya delay karke queue mein daal diya jata hai.
- Common algorithms: **Token Bucket**, **Leaky Bucket**, ya **Fixed Window/Sliding Window**.

#### **Real-world analogy**:
Maan le tu ek toll booth pe kaam karta hai. Ek minute mein sirf 10 gaadiyan hi toll se guzar sakti hain (limit). Agar 15 gaadiyan aati hain, toh baki 5 ko wait karna padega ya reject kar diya jayega. Ye hai rate limiting.

#### **Use Cases**:
- API rate limiting (jaise Twitter/X API mein ek user ek ghante mein max 1000 requests bhej sakta hai).
- Database ya server ko overload se bachane ke liye.
- DDOS attacks ko prevent karne ke liye.

#### **Code Example (Go mein - Token Bucket style)**:
```go
package main

import (
    "fmt"
    "time"
)

type RateLimiter struct {
    tokens   int
    maxTokens int
    refillRate time.Duration
}

func NewRateLimiter(maxTokens int, refillRate time.Duration) *RateLimiter {
    rl := &RateLimiter{
        tokens:     maxTokens,
        maxTokens:  maxTokens,
        refillRate: refillRate,
    }
    go rl.refill()
    return rl
}

func (rl *RateLimiter) refill() {
    ticker := time.NewTicker(rl.refillRate)
    for range ticker.C {
        if rl.tokens < rl.maxTokens {
            rl.tokens++
            fmt.Printf("Token refilled, current tokens: %d\n", rl.tokens)
        }
    }
}

func (rl *RateLimiter) Allow() bool {
    if rl.tokens > 0 {
        rl.tokens--
        return true
    }
    return false
}

func main() {
    // 5 tokens max, har second mein ek token refill
    rl := NewRateLimiter(5, time.Second)

    // Simulate requests
    for i := 1; i <= 7; i++ {
        if rl.Allow() {
            fmt.Printf("Request %d allowed\n", i)
        } else {
            fmt.Printf("Request %d rejected\n", i)
        }
        time.Sleep(200 * time.Millisecond)
    }
}
```
**Explanation**: Yahan `RateLimiter` ek token bucket banata hai. 5 tokens initially hain, aur har second mein ek token refill hota hai. Agar token hai to request allow hoti hai, nahi to reject. Ye simple rate limiting ka example hai.

---

### **Scheduler**
**Scheduler** ka kaam hota hai tasks ko plan aur execute karna, specific time pe ya specific conditions ke hisaab se. Ye ensure karta hai ki tasks sahi waqt pe aur sahi tarike se run hon.

#### **Kaise kaam karta hai?**
- Scheduler tasks ko queue ya list mein rakhta hai.
- Har task ke liye ek time ya condition hoti hai (jaise "har 5 minute mein ye task chala do").
- Scheduler decide karta hai kab aur kaise task run hoga, based on priority, time, ya resource availability.
- Common types: **Cron Scheduler**, **Task Queue Scheduler**, **Real-time Scheduler**.

#### **Real-world analogy**:
Maan le tu ek event organizer hai. Tere paas kai events hain (tasks), jaise "12 baje meeting", "3 baje party", "5 baje speech". Tu ek timetable banata hai aur ensure karta hai ki har event apne time pe ho. Ye kaam scheduler karta hai.

#### **Use Cases**:
- Cron jobs (jaise har raat 2 baje database backup karna).
- Background tasks (jaise email notifications bhejna).
- Distributed systems mein tasks ko nodes pe schedule karna.

#### **Code Example (Go mein - Simple Cron-like Scheduler)**:
```go
package main

import (
    "fmt"
    "time"
)

type Task struct {
    id       int
    name     string
    interval time.Duration
}

func (t *Task) run() {
    fmt.Printf("Running task %s (ID: %d) at %s\n", t.name, t.id, time.Now().Format(time.RFC3339))
}

func scheduler(tasks []Task) {
    for _, task := range tasks {
        go func(t Task) {
            ticker := time.NewTicker(t.interval)
            for range ticker.C {
                t.run()
            }
        }(task)
    }
}

func main() {
    tasks := []Task{
        {id: 1, name: "Send Email", interval: 2 * time.Second},
        {id: 2, name: "Backup DB", interval: 5 * time.Second},
    }

    scheduler(tasks)

    // Run for some time to see the output
    time.Sleep(15 * time.Second)
}
```
**Explanation**: Yahan ek simple scheduler hai jo do tasks ko schedule karta hai:
- "Send Email" task har 2 second mein chalta hai.
- "Backup DB" task har 5 second mein chalta hai.
Har task apne interval ke hisaab se background mein chalta hai.

---

### **Rate Limiter vs Scheduler**
| **Aspect**         | **Rate Limiter**                              | **Scheduler**                              |
|---------------------|----------------------------------------------|-------------------------------------------|
| **Purpose**         | Requests ya tasks ke rate ko control karta hai | Tasks ko specific time ya condition pe run karta hai |
| **Focus**           | Resource usage ko limit karna                 | Tasks ko organize aur execute karna        |
| **Example**         | API calls ko 100/second tak limit karna       | Har 5 minute mein ek report generate karna |
| **Algorithm**       | Token Bucket, Leaky Bucket                   | Cron, Priority Queue, Round Robin          |
| **Use Case**        | Server overload se bachana                    | Background jobs ko automate karna          |

---

### **Key Points**:
1. **Rate Limiter**: System ke resources ko protect karta hai by limiting how often tasks ya requests execute hote hain.
2. **Scheduler**: Tasks ko sahi waqt pe ya sahi sequence mein run karta hai.
3. **Combine Use**: Real-world mein dono sath mein use hote hain. Jaise, ek scheduler har 5 minute mein ek task run karta hai, aur rate limiter ensure karta hai ki wo task system ko overload na kare.

