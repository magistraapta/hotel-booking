package config

import (
	"backend/internal/domain"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedDatabase populates the database with sample hotels, rooms, and facilities
func SeedDatabase(db *gorm.DB) {
	log.Println("Seeding database...")

	// Check if data already exists
	var hotelCount int64
	db.Model(&domain.Hotel{}).Count(&hotelCount)
	if hotelCount > 0 {
		log.Println("Database already seeded. Skipping seed operation.")
		return
	}

	// Create facilities
	facilities := []*domain.Facility{
		{Id: uuid.New(), Name: "WiFi"},
		{Id: uuid.New(), Name: "Air Conditioning"},
		{Id: uuid.New(), Name: "TV"},
		{Id: uuid.New(), Name: "Mini Bar"},
		{Id: uuid.New(), Name: "Room Service"},
		{Id: uuid.New(), Name: "Swimming Pool"},
		{Id: uuid.New(), Name: "Gym"},
		{Id: uuid.New(), Name: "Spa"},
		{Id: uuid.New(), Name: "Parking"},
		{Id: uuid.New(), Name: "Pet Friendly"},
	}

	for _, facility := range facilities {
		if err := db.Create(facility).Error; err != nil {
			log.Printf("Error creating facility %s: %v", facility.Name, err)
		}
	}
	log.Printf("Created %d facilities", len(facilities))

	// Create hotels with rooms
	hotels := []struct {
		hotel domain.Hotel
		rooms []struct {
			size          int
			price         float64
			description   string
			facilityNames []string
		}
	}{
		{
			hotel: domain.Hotel{
				Id:          uuid.New(),
				Name:        "Grand Plaza Hotel",
				Description: "Luxurious 5-star hotel in the heart of the city with stunning views and world-class amenities.",
				Address:     "123 Main Street, Downtown, City 12345",
				Rating:      4.8,
			},
			rooms: []struct {
				size          int
				price         float64
				description   string
				facilityNames []string
			}{
				{size: 25, price: 150.00, description: "Cozy single room with city view", facilityNames: []string{"WiFi", "Air Conditioning", "TV"}},
				{size: 35, price: 220.00, description: "Comfortable double room with balcony", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Mini Bar"}},
				{size: 50, price: 350.00, description: "Spacious suite with living area", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Mini Bar", "Room Service"}},
			},
		},
		{
			hotel: domain.Hotel{
				Id:          uuid.New(),
				Name:        "Oceanview Resort",
				Description: "Beachfront resort offering direct access to pristine beaches and tropical paradise experience.",
				Address:     "456 Beach Boulevard, Coastal Area, City 67890",
				Rating:      4.6,
			},
			rooms: []struct {
				size          int
				price         float64
				description   string
				facilityNames []string
			}{
				{size: 30, price: 180.00, description: "Standard room with ocean view", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Swimming Pool"}},
				{size: 45, price: 280.00, description: "Deluxe room with private balcony", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Mini Bar", "Swimming Pool", "Gym"}},
				{size: 60, price: 450.00, description: "Premium suite with jacuzzi", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Mini Bar", "Room Service", "Swimming Pool", "Gym", "Spa"}},
			},
		},
		{
			hotel: domain.Hotel{
				Id:          uuid.New(),
				Name:        "Mountain Lodge",
				Description: "Rustic mountain retreat perfect for nature lovers and adventure seekers.",
				Address:     "789 Mountain Trail, Highland Valley, City 11111",
				Rating:      4.4,
			},
			rooms: []struct {
				size          int
				price         float64
				description   string
				facilityNames []string
			}{
				{size: 20, price: 120.00, description: "Basic cabin room", facilityNames: []string{"WiFi", "TV", "Parking"}},
				{size: 35, price: 200.00, description: "Family room with mountain view", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Parking", "Pet Friendly"}},
				{size: 40, price: 300.00, description: "Luxury cabin with fireplace", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Mini Bar", "Parking", "Pet Friendly"}},
			},
		},
		{
			hotel: domain.Hotel{
				Id:          uuid.New(),
				Name:        "Business Center Hotel",
				Description: "Modern hotel designed for business travelers with conference facilities and high-speed internet.",
				Address:     "321 Corporate Avenue, Business District, City 22222",
				Rating:      4.5,
			},
			rooms: []struct {
				size          int
				price         float64
				description   string
				facilityNames []string
			}{
				{size: 28, price: 160.00, description: "Standard business room", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Gym"}},
				{size: 38, price: 240.00, description: "Executive room with work desk", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Mini Bar", "Room Service", "Gym"}},
				{size: 55, price: 400.00, description: "Presidential suite with meeting room", facilityNames: []string{"WiFi", "Air Conditioning", "TV", "Mini Bar", "Room Service", "Gym", "Parking"}},
			},
		},
	}

	// Create a map for quick facility lookup
	facilityMap := make(map[string]*domain.Facility)
	for _, facility := range facilities {
		facilityMap[facility.Name] = facility
	}

	// Create hotels and their rooms
	for _, hotelData := range hotels {
		// Create hotel
		if err := db.Create(&hotelData.hotel).Error; err != nil {
			log.Printf("Error creating hotel %s: %v", hotelData.hotel.Name, err)
			continue
		}

		// Create rooms for this hotel
		for _, roomData := range hotelData.rooms {
			room := &domain.Room{
				Id:          uuid.New(),
				Size:        roomData.size,
				Price:       roomData.price,
				Description: roomData.description,
				Available:   true,
				HotelId:     hotelData.hotel.Id,
			}

			// Associate facilities with room
			for _, facilityName := range roomData.facilityNames {
				if facility, exists := facilityMap[facilityName]; exists {
					room.Facilities = append(room.Facilities, facility)
				}
			}

			if err := db.Create(room).Error; err != nil {
				log.Printf("Error creating room for hotel %s: %v", hotelData.hotel.Name, err)
				continue
			}
		}

		log.Printf("Created hotel: %s with %d rooms", hotelData.hotel.Name, len(hotelData.rooms))
	}

	log.Println("Database seeding completed successfully!")
}
