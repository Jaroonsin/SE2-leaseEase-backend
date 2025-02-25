-- Function to update the average rating and review count of a property
CREATE OR REPLACE FUNCTION update_property_rating_and_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE properties
    SET 
        avg_rating = (
            SELECT COALESCE(AVG(reviews.rating), 0)
            FROM property_reviews
            JOIN reviews ON property_reviews.review_id = reviews.id
            WHERE property_reviews.property_id = NEW.property_id
        ),
        review_count = (
            SELECT COUNT(*)
            FROM property_reviews
            WHERE property_reviews.property_id = NEW.property_id
        )
    WHERE id = NEW.property_id;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
