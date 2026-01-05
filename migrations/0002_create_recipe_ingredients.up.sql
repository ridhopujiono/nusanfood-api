CREATE TABLE IF NOT EXISTS recipe_ingredients (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    recipe_id BIGINT UNSIGNED NOT NULL,
    food_id BIGINT UNSIGNED NOT NULL,
    food_serving_id BIGINT UNSIGNED NOT NULL,
    quantity DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_recipe_id (recipe_id),
    INDEX idx_food_id (food_id),
    INDEX idx_food_serving_id (food_serving_id),

    CONSTRAINT fk_recipe_ingredients_recipe
        FOREIGN KEY (recipe_id) REFERENCES recipes(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_recipe_ingredients_food
        FOREIGN KEY (food_id) REFERENCES foods(id),

    CONSTRAINT fk_recipe_ingredients_serving
        FOREIGN KEY (food_serving_id) REFERENCES food_servings(id)
) ENGINE=InnoDB;
