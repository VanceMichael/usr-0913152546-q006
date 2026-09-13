package main
import("database/sql";"log";"net/http";_ "modernc.org/sqlite")
func main(){db,e:=sql.Open("sqlite","data.db");if e!=nil{log.Fatal(e)};db.Exec("create table if not exists jobs(id integer primary key,status text)");http.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){w.Write([]byte(`{"status":"ok"}`))});log.Fatal(http.ListenAndServe(":8080",nil))}
