package models

type Participant struct {
    ID    uint   `gorm:"primary_key" json:"id"`
    Username string   `json:"username"`
    Events []Event `gorm:"many2many:event_participants;" json:"events"` // Many-to-Many relationship
    // liked  events...
    // liked promoters...

    // subscribed promoters (notifiche su nuovi eventi)
    Friends []Participant `gorm:"many2many:participant_friends;" json:"friends"`
    IsPrivate bool 

}


