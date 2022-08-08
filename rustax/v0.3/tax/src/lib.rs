mod tax;

use std::fmt;
use crate::tax::*;

// Calculator is a simple structure to calculate taxes on an income
#[derive(Debug)]
pub struct Calculator {
    pub income: f32,

    tax_allowance: f32,

    taxes: [Tax; 5],
}

impl Calculator {
    // New create a new instance of Tax
    pub fn new(income: f32, tax_allowance: f32) -> Self {
        Self {
            income,
            tax_allowance,
            taxes: [
                Tax::new(0.0, 10225.0, 0.0),
                Tax::new(10225.0, 26070.0, 11.0),
                Tax::new(26070.0, 74545.0, 30.0),
                Tax::new(74545.0, 160336.0, 41.0),
                Tax::new(160337.0, f32::MAX, 45.0),
            ],
        }
    }

    pub fn tax(&self) -> f32 {
        let mut taxes: f32 = 0.0;
        let income = self.income / 100.0 * (100.0 - self.tax_allowance);

        for t in self.taxes {
            if income > t.max {
                taxes += (t.max - t.min) / 100.0 * t.percentage
            } else {
                taxes += (income - t.min) / 100.0 * t.percentage;
                break;
            }
        }
        
        return taxes
    }
}

impl fmt::Display for Calculator {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({}, {}, {:#?})", self.income, self.tax_allowance, self.taxes)
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn can_create() {
        let calculator = crate::Calculator::new(4.0, 0.0);
        assert_eq!(calculator.income, 4.0)
    }

    // Based on https://www.service-public.fr/particuliers/vosdroits/F1419
    #[test]
    fn can_compute() {
        let calculator = crate::Calculator::new(30_000.0, 0.0);
        assert_eq!(calculator.tax(), 2_921.95)
    }

    #[test]
    fn can_compute_with_tax_allowance() {
        let calculator = crate::Calculator::new(30_000.0, 30.0);
        assert_eq!(calculator.tax(), 1_185.25)
    }
}
