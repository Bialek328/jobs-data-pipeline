package main

// func InsertJobListingIntoDB(db *sql.DB, job JobOffer) error {
//     insertQuery := `INSERT INTO job_listings (url, tech_stack, misc_info, salary)
//                     VALUES ($1, $2, $3, $4);`
//     _, err := db.Exec(insertQuery, job.Url, job.TechStack, job.MiscInfo, job.Salary)
//     if err != nil {
//         return err
//     }
//     return nil
// }
