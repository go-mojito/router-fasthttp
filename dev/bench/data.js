window.BENCHMARK_DATA = {
  "lastUpdate": 1651700921904,
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
      },
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
          "id": "f9d61880ff2e992c1fd70f306aa2e4aa85f49877",
          "message": "fix(test): Wait 2 seconds for server",
          "timestamp": "2022-05-04T23:31:33+02:00",
          "tree_id": "81db0537a59db069169be594f71a658650b4dbda",
          "url": "https://github.com/go-mojito/router-fasthttp/commit/f9d61880ff2e992c1fd70f306aa2e4aa85f49877"
        },
        "date": 1651699954698,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Router_Handler",
            "value": 328115,
            "unit": "ns/op\t   30719 B/op\t     214 allocs/op",
            "extra": "3320 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_Not_Found",
            "value": 328939,
            "unit": "ns/op\t   29400 B/op\t     181 allocs/op",
            "extra": "3735 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_With_Middleware",
            "value": 309087,
            "unit": "ns/op\t   29409 B/op\t     181 allocs/op",
            "extra": "3380 times\n2 procs"
          }
        ]
      },
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
          "id": "f40efe3e01727ef61445417d1a47ee0e147ea52f",
          "message": "fix(test): Fix default test",
          "timestamp": "2022-05-04T23:47:44+02:00",
          "tree_id": "bdb3df70417750f7c567c57b1236b55bdf221971",
          "url": "https://github.com/go-mojito/router-fasthttp/commit/f40efe3e01727ef61445417d1a47ee0e147ea52f"
        },
        "date": 1651700920687,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Router_Handler",
            "value": 581478,
            "unit": "ns/op\t   31181 B/op\t     216 allocs/op",
            "extra": "2064 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_Not_Found",
            "value": 493246,
            "unit": "ns/op\t   29346 B/op\t     181 allocs/op",
            "extra": "2256 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_With_Middleware",
            "value": 477278,
            "unit": "ns/op\t   29444 B/op\t     181 allocs/op",
            "extra": "2635 times\n2 procs"
          }
        ]
      }
    ]
  }
}