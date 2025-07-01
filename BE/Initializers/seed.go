package Initializers

import (
	"fmt"
	"time"
	"wan-api-kol-event/Models"
)

func SeedKols() {
	for i := 1; i <= 20; i++ {
		kol := Models.Kol{
			KolID:                int64(1000 + i),
			UserProfileID:        int64(2000 + i),
			Language:             "en",
			Education:            "Bachelor's in Computer Science",
			ExpectedSalary:       50000 + int64(i*1000),
			ExpectedSalaryEnable: true,
			ChannelSettingTypeID: 1,
			IDFrontURL:           "https://example.com/id-front.jpg",
			IDBackURL:            "https://example.com/id-back.jpg",
			PortraitURL:          "https://example.com/portrait.jpg",
			RewardID:             300 + int64(i),
			PaymentMethodID:      400 + int64(i),
			TestimonialsID:       500 + int64(i),
			VerificationStatus:   true,
			Enabled:              true,
			ActiveDate:           time.Now(),
			Active:               true,
			CreatedBy:            "admin",
			CreatedDate:          time.Now(),
			ModifiedBy:           "admin",
			ModifiedDate:         time.Now(),
			IsRemove:             false,
			IsOnBoarding:         true,
			Code:                 fmt.Sprintf("KOL20240%03d", i),
			PortraitRightURL:     "https://example.com/portrait-right.jpg",
			PortraitLeftURL:      "https://example.com/portrait-left.jpg",
			LivenessStatus:       true,
		}

		DB.Create(&kol)
	}

	fmt.Println("✅ Seeded 20 KOL records successfully.")
}
