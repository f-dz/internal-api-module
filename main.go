package main

import (
	"context"
	"fmt"

	"github.com/api-abc/internal-api-module/data"
	"github.com/api-abc/internal-api-module/model/domain"
	"github.com/api-abc/internal-api-module/model/request"
	"github.com/api-abc/internal-api-module/rest"
)

func main() {
	// =============================================== //
	// TEST INSERT
	req := request.InsertRequest{
		Name: "Dipsy",
		Age:  20,
		JobDetails: domain.Job{
			Position:            "manager",
			YearsWorkExperience: 5,
			WorkStatus:          "work",
		},
	}

	cl := rest.New("http://localhost:8080")
	cl_data := data.New(cl)

	data_insert, err := cl_data.Insert(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data_insert)
	}

	countData, err := cl_data.GetInserted(context.Background())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(countData)
	}

	// =============================================== //
	// TEST UPDATE
	req2 := request.UpdateRequest{
		Age: 60,
		JobDetails: domain.Job{
			Position:            "employee",
			YearsWorkExperience: 2,
			WorkStatus:          "work",
		},
	}

	cl2 := rest.New("http://localhost:8082")
	cl2_data := data.New(cl2)

	data_update, err2 := cl2_data.Update(context.Background(), req2, "Dipsy")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(data_update)
	}

	countData2, err2 := cl_data.GetUpdated(context.Background())
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(countData2)
	}

	// =============================================== //
	// TEST DELETE
	cl3 := rest.New("http://localhost:8081")
	cl3_data := data.New(cl3)

	data_delete, err3 := cl3_data.Delete(context.Background(), "Dipsy")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(data_delete)
	}

	countData3, err3 := cl_data.GetDeleted(context.Background())
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(countData3)
	}
}
