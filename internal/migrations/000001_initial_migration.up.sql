CREATE TYPE reg_state_enum as ENUM ('state_agreements', 'state_name', 'state_sex', 'state_main_picture', 'state_another_picture', 'state_target');
CREATE TYPE sex_enum as ENUM ('Male', 'Female', 'NotSelected');

CREATE TYPE target_name_enum AS ENUM ('for_like', 'for_one_day', 'for_all_time');

CREATE TABLE targets (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name target_name_enum NOT_NULL
);

CREATE TYPE category_name_enum AS ENUM ('Interests', 'Worldview', 'Psychographer', 'Activity');

CREATE TABLE category (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name category_name_enum,
    volume INTEGER,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
);

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(255) UNIQUE,
  registration_state reg_state_enum,
  name VARCHAR(50),
  bio VARCHAR(500),
  birthdate timestamptz,
  sex sex_enum, 
  location VARCHAR(255),
  height INTEGER,
  target_id UUID,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),

  CONSTRAINT fk_targets FOREIGN KEY (target_id) REFERENCES targets(id)
);

CREATE TABLE pictures (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  is_main BOOLEAN DEFAULT false,
  approved BOOLEAN DEFAULT false,
  path VARCHAR(255),

  CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_pictures_user_id ON pictures(user_id);

-- У юзера может быть только одна главная фотка (is_main = true).
-- Этот индекс не даст вставить вторую true, но разрешит много false.
CREATE UNIQUE INDEX idx_pictures_one_main_per_user 
ON pictures (user_id) 
WHERE is_main IS TRUE;


CREATE TABLE tags (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255),
  volume INTEGER,
  category_id UUID, 
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),

  CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE INDEX idx_tags_category_id ON tags(category_id);
CREATE INDEX idx_tags_name ON tags(name);

CREATE TABLE user_tags (
    user_id UUID NOT NULL,
    tag_id UUID NOT NULL,
    PRIMARY KEY (user_id, tag_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE, 
    CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

CREATE INDEX idx_user_tags_tag_id ON user_tags(tag_id);