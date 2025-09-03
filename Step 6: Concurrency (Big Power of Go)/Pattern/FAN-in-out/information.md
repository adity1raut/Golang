Bhai, concurrency mein **Fan-In** aur **Fan-Out** patterns ka concept samajhna bilkul easy hai agar hum ise simple language mein tod ke dekhein. Ye dono patterns Go (Golang) ya aur concurrent programming mein commonly use hote hain, especially jab **goroutines** aur **channels** ke sath kaam karte hain. Chalo, ek ek karke samajhte hain:

---

### **Fan-Out Pattern**
**Fan-Out** ka matlab hai ek single input ko multiple workers (goroutines) ke beech baant dena, taaki kaam parallel mein ho sake. Yani, ek bada task multiple chhote tasks mein split hota hai, aur har task ko alag-alag worker handle karta hai.

#### **Kaise kaam karta hai?**
- Ek single input source hota hai (jaise ek channel ya queue).
- Is input ko multiple workers (goroutines) ke beech distribute kar diya jata hai.
- Har worker apna hissa ka kaam independently process karta hai.

#### **Example (real-world analogy):**
Maan le tu ek restaurant mein chef hai, aur ek bada order aata hai jisme 100 rotis banana hai. Tu akela 100 rotis nahi bana sakta quickly, toh tu 5 aur chefs ko bula leta hai. Har chef ko 20 rotis banane ka kaam de diya. Ye hai **Fan-Out** – ek bada task (100 rotis) ko multiple workers (chefs) mein baant diya.

#### **Code Example (Go mein):**
```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
    }
}

func main() {
    jobs := make(chan int, 100)
    var wg sync.WaitGroup

    // 5 workers start karo (Fan-Out)
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, jobs, &wg)
    }

    // Jobs bhejo
    for j := 1; j <= 20; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
}
```
**Explanation**: Yahan ek `jobs` channel hai jisme tasks (1 se 20) bheje ja rahe hain. 5 workers (goroutines) hain jo in tasks ko parallel mein process karte hain. Ye **Fan-Out** hai kyunki ek input channel se tasks multiple workers mein distribute ho rahe hain.

---

### **Fan-In Pattern**
**Fan-In** ka matlab hai multiple input sources (channels) se data ko ek single output channel mein combine karna. Yani, kai workers apna-apna output ek jagah collect karte hain.

#### **Kaise kaam karta hai?**
- Multiple workers (goroutines) apna kaam karke apne results ko alag-alag channels pe bhejte hain.
- Ek single channel hota hai jo in sabhi workers ke outputs ko collect karta hai.
- Ye pattern tab useful hai jab multiple sources se data ko ek jagah merge karna ho.

#### **Example (real-world analogy):**
Wapas restaurant ka scene lete hain. Ab maan le 5 chefs ne apne-apne hisse ki rotis bana li hain (20-20 rotis). Ab ek waiter hai jo in sab chefs ke banaye rotis ko ek bade tray mein collect karta hai aur customer tak le jata hai. Ye collect karne ka process hai **Fan-In** – multiple sources (chefs) se output ek jagah (waiter ka tray) pe aata hai.

#### **Code Example (Go mein):**
```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 3; i++ {
        results <- fmt.Sprintf("Result from worker %d: %d", id, i)
    }
}

func main() {
    results := make(chan string, 100)
    var wg sync.WaitGroup

    // 3 workers start karo
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, results, &wg)
    }

    // Fan-In: Results ko ek channel mein collect karo
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect kiye hue results ko print karo
    for result := range results {
        fmt.Println(result)
    }
}
```
**Explanation**: Yahan 3 workers apne results ko `results` channel mein bhejte hain. Ek single channel (`results`) sab workers ke outputs ko collect karta hai. Ye hai **Fan-In**.

---

### **Fan-In aur Fan-Out Combine**
Aksar real-world applications mein **Fan-Out** aur **Fan-In** ek sath use hote hain. For example:
- **Fan-Out**: Ek bada task multiple workers mein baant do.
- **Fan-In**: Un workers ke results ko ek jagah collect karo.

#### **Real-world analogy**:
Restaurant mein 100 rotis ka order aaya (input). 5 chefs (workers) ko kaam baant diya (Fan-Out). Har chef apni rotis banata hai, aur ek waiter un sabki rotis ko ek tray mein collect karta hai (Fan-In) aur customer ko deta hai.

#### **Code Example (Combined):**
```go
package main

import (
    "fmt"
    "sync"
)

func producer(jobs chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 10; i++ {
        jobs <- i
    }
}

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        results <- fmt.Sprintf("Worker %d processed job %d", id, job)
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan string, 100)
    var wgProducers, wgWorkers sync.WaitGroup

    // Fan-Out: 2 producers tasks generate karte hain
    wgProducers.Add(2)
    go producer(jobs, &wgProducers)
    go producer(jobs, &wgProducers)

    // Fan-Out: 3 workers jobs process karte hain
    wgWorkers.Add(3)
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results, &wgWorkers)
    }

    // Producers ke complete hone ka wait aur jobs channel band karo
    go func() {
        wgProducers.Wait()
        close(jobs)
    }()

    // Workers ke complete hone ka wait aur results channel band karo
    go func() {
        wgWorkers.Wait()
        close(results)
    }()

    // Fan-In: Results ko collect aur print karo
    for result := range results {
        fmt.Println(result)
    }
}
```
**Explanation**: 
- 2 producers tasks generate karte hain (Fan-Out for input).
- 3 workers in tasks ko process karte hain (Fan-Out for processing).
- Ek `results` channel sab workers ke output ko collect karta hai (Fan-In).

---

### **Key Points**:
1. **Fan-Out**: Ek input source se multiple workers ko kaam baantna.
2. **Fan-In**: Multiple workers ke outputs ko ek channel mein merge karna.
3. **Use Case**: Ye patterns scalable systems mein use hote hain, jaise data processing pipelines, distributed systems, ya microservices.
4. **Go mein**: Channels aur goroutines in patterns ko implement karne ka sabse common tareeka hain.

