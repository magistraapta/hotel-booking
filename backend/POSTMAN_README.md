# Postman Collection - Hotel Booking API

## Import Instructions

1. Open Postman
2. Click **Import** button (top left)
3. Select the file: `Hotel_Booking_API.postman_collection.json`
4. The collection will be imported with all endpoints

## Environment Setup

The collection uses a variable `{{base_url}}` which defaults to `http://localhost:8080`.

### To change the base URL:
1. Click on the collection name
2. Go to **Variables** tab
3. Update `base_url` value (e.g., `http://localhost:8080` or your production URL)

## API Endpoints

### Users

#### POST /users/
Create a new user account

**Request Body:**
```json
{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123",
    "is_admin": false
}
```

**Response (201 Created):**
```json
{
    "message": "User created successfully"
}
```

---

### Hotels

#### POST /hotels/
Create a new hotel

**Request Body:**
```json
{
    "name": "Grand Plaza Hotel",
    "description": "Luxurious 5-star hotel in the heart of the city",
    "address": "123 Main Street, Downtown, City 12345",
    "rating": 4.8
}
```

**Response (201 Created):**
```json
{
    "message": "Hotel created successfully"
}
```

#### GET /hotels/
Get all hotels with rooms and facilities

**Response (200 OK):**
```json
{
    "message": "Hotels fetched successfully",
    "hotels": [
        {
            "id": "uuid",
            "name": "Grand Plaza Hotel",
            "description": "...",
            "address": "...",
            "rating": 4.8,
            "rooms": [...]
        }
    ]
}
```

#### GET /hotels/:id
Get a specific hotel by ID

**Path Parameters:**
- `id` (UUID): Hotel ID

**Response (200 OK):**
```json
{
    "message": "Hotel fetched successfully",
    "hotel": {
        "id": "uuid",
        "name": "Grand Plaza Hotel",
        "description": "...",
        "address": "...",
        "rating": 4.8,
        "rooms": [...]
    }
}
```

---

### Bookings

#### POST /bookings/
Create a new booking

**Request Body:**
```json
{
    "user_id": "user-uuid-here",
    "hotel_id": "hotel-uuid-here",
    "room_id": "room-uuid-here",
    "check_in_date": "2025-12-20T00:00:00Z",
    "check_out_date": "2025-12-23T00:00:00Z",
    "total_price": 450.00,
    "is_cancelled": false
}
```

**Response (201 Created):**
```json
{
    "message": "Booking created successfully"
}
```

#### GET /bookings/
Get all bookings

**Response (200 OK):**
```json
{
    "message": "Bookings fetched successfully",
    "bookings": [
        {
            "id": "uuid",
            "user_id": "uuid",
            "hotel_id": "uuid",
            "room_id": "uuid",
            "check_in_date": "2025-12-20T00:00:00Z",
            "check_out_date": "2025-12-23T00:00:00Z",
            "total_price": 450.00,
            "is_cancelled": false
        }
    ]
}
```

#### GET /bookings/:id
Get a specific booking by ID

**Path Parameters:**
- `id` (UUID): Booking ID

**Response (200 OK):**
```json
{
    "message": "Booking fetched successfully",
    "booking": {
        "id": "uuid",
        "user_id": "uuid",
        "hotel_id": "uuid",
        "room_id": "uuid",
        "check_in_date": "2025-12-20T00:00:00Z",
        "check_out_date": "2025-12-23T00:00:00Z",
        "total_price": 450.00,
        "is_cancelled": false
    }
}
```

## Error Responses

All endpoints return error responses in the following format:

**Status Code: 500 Internal Server Error**
```json
{
    "error": "Error message here"
}
```

## Testing Workflow

1. **Start the server:**
   ```bash
   cd backend/cmd
   go run main.go
   ```

2. **Get seeded data:**
   - The database is automatically seeded with sample hotels, rooms, and facilities
   - Use `GET /hotels/` to get hotel IDs and room IDs for testing bookings

3. **Create a user:**
   - Use `POST /users/` to create a test user
   - Save the user ID from the response (if returned) or use a UUID

4. **Create a booking:**
   - Use hotel and room IDs from the seeded data
   - Use the user ID from step 3
   - Use `POST /bookings/` to create a booking

## Notes

- All IDs are UUIDs (format: `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`)
- Dates should be in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
- The server runs on port 8080 by default (check your `.env` file for `SERVER_PORT`)
- Make sure your database is running and connected before testing

