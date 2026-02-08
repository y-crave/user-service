CREATE TYPE reg_state_enum as ENUM ('state_agreements', 'state_name', 'state_sex', 'state_main_picture', 'state_another_picture', 'state_target');
CREATE TYPE sex_enum as ENUM ('Male', 'Female', 'Not selected');

CREATE TYPE target_name_enum AS ENUM ('for_like', 'for_one_day', 'for_all_time');

CREATE TABLE targets (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name target_name_enum
);

CREATE TYPE category_name_enum AS ENUM ('Интересы', 'Мировоззрение', 'Психограф', 'Активность');

CREATE TABLE category (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name category_name_enum,
    volume INTEGER,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(50),
  registration_state reg_state_enum,
  name VARCHAR(20),
  bio VARCHAR(500),
  birthdate timestamptz,
  sex sex_enum, 
  location VARCHAR(255),
  height INTEGER,
  target_id UUID,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz,

  CONSTRAINT fk_targets FOREIGN KEY (target_id) REFERENCES targets(id)
);

CREATE TABLE pictures (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID,
  is_main BOOLEAN,
  approved BOOLEAN,
  path VARCHAR(225),

  CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE tags (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255),
  volume INTEGER,
  category_id UUID, 
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz,

  CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE user_tags (
    user_id UUID NOT NULL,
    tag_id UUID NOT NULL,
    PRIMARY KEY (user_id, tag_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE, 
    CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

