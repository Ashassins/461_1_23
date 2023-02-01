// url/repo.go
package url
import (
    "errors"

    "github.com/jinzhu/gorm"
)

type URLRepository struct {
    database *gorm.DB
}

func (repository *URLRepository) FindAll() []URL {
    var urls []URL
    repository.database.Find(&urls)
    return urls
}

func (repository *URLRepository) Find(id int) (URL, error) {
    var url URL
    err := repository.database.Find(&url, id).Error
    if url.StoredURL == "" {
        err = errors.New("URL not found")
    }
    return url, err
}

func (repository *URLRepository) Create(url URL) (URL, error) {
    err := repository.database.Create(&url).Error
    if err != nil {
        return url, err
    }

    return url, nil
}

func (repository *URLRepository) Save(user URL) (URL, error) {
    err := repository.database.Save(user).Error
    return user, err
}

func (repository *URLRepository) Delete(id int) int64 {
    count := repository.database.Delete(&URL{}, id).RowsAffected
    return count
}

func NewURLRepository(database *gorm.DB) *URLRepository {
    return &URLRepository{
		database: database,
    }
}