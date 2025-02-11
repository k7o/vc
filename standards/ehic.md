# EHIC

```json
{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "type": "object",
    "properties": {
        "subject": {
            "type": "object",
            "properties": {
                "forename": {
                    "type": "string"
                },
                "family_name": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string",
                    "format": "date"
                },
                "other_elements": {
                    "type": "object",
                    "properties": {
                        "sex": {
                            "$ref": "#/$defs/sex_type"
                        },
                        "forename_at_birth": {
                            "type": "string"
                        },
                        "family_name_at_birth": {
                            "type": "string"
                        }
                    }
                }
            },
            "required": [
                "forename",
                "family_name",
                "date_of_birth"
            ]
        },
        "social_security_pin": {
            "type": "string",
            "description": "Personal Identification Number as defined in the issuing institution"
        },
        "period_entitlement": {
            "type": "object",
            "properties": {
                "starting_date": {
                    "type": "string",
                    "format": "date"
                },
                "ending_date": {
                    "type": "string",
                    "format": "date"
                }
            },
            "required": [
                "starting_date",
                "ending_date"
            ]
        },
        "document_id": {
            "type": "string",
            "pattern": "^[-a-zA-Z0-9]{1,65}$"
        },
        "competent_institution": {
            "type": "object",
            "properties": {
                "institution_id": {
                    "$ref": "#/$defs/eessi_institution_id"
                },
                "institution_name": {
                    "type": "string"
                },
                "institution_country": {
                    "$ref": "#/$defs/iso3166_1_eu_efta_country_code"
                }
            },
            "required": [
                "institution_id",
                "institution_country"
            ]
        }
    },
    "required": [
        "social_security_pin",
        "period_entitlement",
        "document_id",
        "competent_institution"
    ],
    "$defs": {
        "iso3166_1_eu_efta_country_code": {
            "type": "string",
            "pattern": "^(AT|BE|BG|HR|CY|CZ|DK|EE|FI|FR|DE|EL|HU|IS|IE|IT|LV|LI|LT|LU|MT|NL|NO|PL|PT|RO|SK|SI|ES|SE|CH|UK|EU){1}$",
            "description": "Country code according to EU/EFTA-Countries according to ISO-3166-1 + UK"
        },
        "sex_type": {
            "type": "string",
            "pattern": "^(01|02|98){1}$",
            "description": "01 - Male, 02 - Female, 98 - Unknown"
        },
        "eessi_institution_id": {
            "type": "string",
            "pattern": "^(AT|BE|BG|HR|CY|CZ|DK|EE|FI|FR|DE|EL|HU|IS|IE|IT|LV|LI|LT|LU|MT|NL|NO|PL|PT|RO|SK|SI|ES|SE|CH|UK|EU):[a-zA-Z0-9]{4,10}$",
            "description": "Institution ID in the format 'AT:19789'"
        }
    }
}
```

## Example

![EHIC Expample](./img/EHIC_example.png)
