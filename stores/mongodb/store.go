package mongostore

// import "log"

// // Store contains a DatabaseHelper
// type Store struct {
// 	db mongodb.DatabaseHelper
// }

// // New createa new new store connected to the given MongoDB connection URi
// func New(connURI string) (*Store, error) {
// 	dbh, err := mongodb.NewDatabaseHelperFromURI(connURI)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Store{db: dbh}, nil
// }

// // NewDatabaseHelperFromURI returns a new DatabaseHelper from a URI
// func NewDatabaseHelperFromURI(connURI string) (DatabaseHelper, error) {
// 	mcs, err := connstring.Parse(connURI)
// 	if err != nil {
// 		log.Printf("Error parsing MongoDB url: %s", err)
// 		return nil, err
// 	}

// 	client, err := NewClient(options.Client().ApplyURI(connURI))
// 	if err != nil {
// 		log.Printf("Error creating MongoDB client: %s", err)
// 		return nil, err
// 	}
// 	err = client.Connect()
// 	if err != nil {
// 		log.Printf("Unable to connect to MongoDB: %s", err)
// 		return nil, err
// 	}
// 	log.Printf("Connected to to MongoDB: %s", mcs.AppName)

// 	db := client.Database(mcs.Database)

// 	return db, nil
// }
