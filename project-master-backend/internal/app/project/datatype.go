package project

import (
    "strconv"
    "strings"
    "time"

    "github.com/jinzhu/now"
    "github.com/gbrlsnchs/jwt"
    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/mongodb/mongo-go-driver/bson/primitive"
)

/*
|--------------------------------------------------------------------------
| Filter Entities
|--------------------------------------------------------------------------
|
*/
type Filter struct {
    Option      string  `json:"option"`
    Operator    string  `json:"operator"`
    Value       string  `json:"value"`
}

func (f *Filter) ToBson() bson.D {
    var filter bson.D

    switch strings.ToUpper(f.Operator) {
        case ">":
            if num, err := strconv.Atoi(f.Value); err == nil {
                filter = bson.D{{f.Option, bson.D{{"$gt", num}}}}
            }
            break
        case "<":
            if num, err := strconv.Atoi(f.Value); err == nil {
                filter = bson.D{{f.Option, bson.D{{"$lt", num}}}}
            }
            break
        case ">=":
            if num, err := strconv.Atoi(f.Value); err == nil {
                filter = bson.D{{f.Option, bson.D{{"$gte", num}}}}
            }
            break
        case "<=":
            if num, err := strconv.Atoi(f.Value); err == nil {
                filter = bson.D{{f.Option, bson.D{{"$lte", num}}}}
            }
            break
        case "LIKE":
            filter = bson.D{{f.Option, primitive.Regex{Pattern: ".*" + f.Value + ".*", Options: "i"}}}
            break
        case "=":
            filter = bson.D{{f.Option, f.Value}}
            break
        case "IN":
            values := strings.Split(f.Value, ",")
            filter = bson.D{{f.Option, bson.M{"$in" : values}}}
            break
        case "DATE":
            if date, err := time.Parse("2006-01-02", f.Value); err == nil {
                filter = bson.D{{f.Option, bson.M{"$gte": now.New(date).BeginningOfDay(), "$lte": now.New(date).EndOfDay()}}}
            }
            break
        case "MONTH":
            if date, err := time.Parse("2006-01-02", f.Value); err == nil {
                filter = bson.D{{f.Option, bson.M{"$gte": now.New(date).BeginningOfMonth(), "$lte": now.New(date).EndOfMonth()}}}
            }
            break
        case "YEAR":
            if date, err := time.Parse("2006-01-02", f.Value); err == nil {
                filter = bson.D{{f.Option, bson.M{"$gte": now.New(date).BeginningOfYear(), "$lte": now.New(date).EndOfYear()}}}
            }
            break
        // default:
        //     filter = bson.D{{f.Option, primitive.Regex{Pattern: f.Value, Options: "i"}}}
        //     break
    }

    return filter
}

/*
|--------------------------------------------------------------------------
| Sort Entities
|--------------------------------------------------------------------------
|
*/
type Sort struct {
    Option      string  `json:"option"`
    Value       string  `json:"value"`
}

func (s *Sort) ToBson() bson.D{
    var sort bson.D

    if s.Option != "" {
        value := 1

        if strings.ToLower(s.Value) == "desc" {
            value = -1
        }

        sort = bson.D{{s.Option, value }}
    } else {
        sort = nil
    }

    return sort
}

/*
|--------------------------------------------------------------------------
| Token Entities
|--------------------------------------------------------------------------
|
*/
type Token struct {
    *jwt.JWT
    ID          string      `json:"sub"`
    Role        []string    `json:"role"`
    FullName    string      `json:"full_name"`
}

func (t *Token) ToUser() AuthUser {
    var authUser AuthUser
    
    authUser.ID         = t.ID
    authUser.FullName   = t.FullName

    return authUser
}

/*
|--------------------------------------------------------------------------
| PageRequest Entities
|--------------------------------------------------------------------------
|
*/
type PageRequest struct {
    Page     int64
    Paginate int64
    PerPage  int64
    Search   string
    Status   string
    Filters  []Filter
    Sorts    []Sort
    Type     string
}