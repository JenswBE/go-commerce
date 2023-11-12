const ibmCloudValidationRules = require("@ibm-cloud/openapi-ruleset");
const {
  propertyCasingConvention,
} = require("@ibm-cloud/openapi-ruleset/src/functions");
const {
  schemas,
} = require("@ibm-cloud/openapi-ruleset-utilities/src/collections");

module.exports = {
  extends: ibmCloudValidationRules,
  rules: {
    "ibm-enum-casing-convention": {
      description: "Enum names must follow uppercase snake convention",
      message: "{{error}}",
      resolved: true,
      given: schemas,
      severity: "warn",
      then: {
        function: propertyCasingConvention,
        functionOptions: {
          type: "macro", // == UPPERCASE_SNAKE
        },
      },
    },
  },
};
