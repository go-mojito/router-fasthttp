window.BENCHMARK_DATA = {
  "lastUpdate": 1651698993120,
  "repoUrl": "https://github.com/go-mojito/router-fasthttp",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "0skillallluck@pm.me",
            "name": "Cedric Lewe",
            "username": "0SkillAllLuck"
          },
          "committer": {
            "email": "0skillallluck@pm.me",
            "name": "Cedric Lewe",
            "username": "0SkillAllLuck"
          },
          "distinct": true,
          "id": "4d98182071069c91d4d7acc9ad64be128eeb166f",
          "message": "fix(ci): CodeQL removed unnessecary checkout",
          "timestamp": "2022-05-04T23:12:14+02:00",
          "tree_id": "4cd77f31c1ff7cc7235f022c91aca04c3c4baa16",
          "url": "https://github.com/go-mojito/router-fasthttp/commit/4d98182071069c91d4d7acc9ad64be128eeb166f"
        },
        "date": 1651698992152,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Router_Handler",
            "value": 283049,
            "unit": "ns/op\t   30915 B/op\t     214 allocs/op",
            "extra": "3867 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_Not_Found",
            "value": 252513,
            "unit": "ns/op\t   29257 B/op\t     181 allocs/op",
            "extra": "4438 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_With_Middleware",
            "value": 726776,
            "unit": "ns/op\t   29830 B/op\t     182 allocs/op",
            "extra": "4444 times\n2 procs"
          }
        ]
      }
    ]
  }
}