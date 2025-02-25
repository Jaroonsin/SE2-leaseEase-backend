-- Trigger to execute function when a new property review is inserted
CREATE OR REPLACE TRIGGER trigger_update_property_rating_and_count
AFTER INSERT ON property_reviews
FOR EACH ROW
EXECUTE FUNCTION update_property_rating_and_count();
