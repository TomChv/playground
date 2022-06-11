use std::fmt;

#[derive(Debug, Clone, Copy)]
pub struct Tax {
    pub min: f32,
    pub max: f32,
    pub percentage: f32, // Value between 1 and 100
}

impl Tax {
    pub fn new(min: f32, max: f32, percentage: f32) -> Self {
        Self {
            min,
            max,
            percentage,
        }
    }
}

impl fmt::Display for Tax {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}, {}, {})", self.min, self.max, self.percentage)
    }
}