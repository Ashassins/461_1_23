// // url/repo.go
package url
// import (
//     "errors"
// )

// type URLRepository struct {
    
// }

// func (repository *URLRepository) FindAll() []URL {
//     return URL
// }
 
// func (repository *URLRepository) Find(id int) (URL, error) {
//     return URL, error
// }

// func (repository *URLRepository) Create(url URL) (URL, error) {
//     return URL, error
// }

// func (repository *URLRepository) Save(user URL) (URL, error) {
//     return URL, error
// }

// func (repository *URLRepository) Delete(id int) int64 {
//     return id;
// }

// // func NewURLRepository(database *gorm.DB) *URLRepository {

// // }

// // Old code from API boilerplate 
// // func (repository *URLRepository) FindAll() []URL {
// //     var urls []URL
// //     // repository..Find(&urls)
// //     return urls
// // }

// // func (repository *URLRepository) Find(id int) (URL, error) {
// //     var url URL
// //     // err := repository..Find(&url, id).Error
// //     if url.StoredURL == "" {
// //         err = errors.New("URL not found")
// //     }
// //     return url, err
// // }

// // func (repository *URLRepository) Create(url URL) (URL, error) {
// //     // err := repository..Create(&url).Error
// //     if err != nil {
// //         return url, err
// //     }

// //     return url, nil
// // }

// // func (repository *URLRepository) Save(user URL) (URL, error) {
// //     // err := repository..Save(user).Error
// //     return user, err
// // }

// // func (repository *URLRepository) Delete(id int) int64 {
// //     // count := repository..Delete(&URL{}, id).RowsAffected
// //     return count
// // }

// // func NewURLRepository( *gorm.DB) *URLRepository {
// //     return &URLRepository{
// // 		// : ,
// //     }
// // }
