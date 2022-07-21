const ibmCloudValidationRules = require("@ibm-cloud/openapi-ruleset");
const {
  enumCaseConvention,
} = require("@ibm-cloud/openapi-ruleset/src/functions");
const { schemas } = require("@ibm-cloud/openapi-ruleset/src/collections");

module.exports = {
  extends: ibmCloudValidationRules,
  rules: {
    "enum-case-convention": {
      // Note 2
      description: "Enum names must follow uppercase snake convention",
      message: "{{error}}",
      resolved: true,
      given: schemas,
      severity: "warn",
      then: {
        function: enumCaseConvention,
        functionOptions: {
          type: "macro", // == UPPERCASE_SNAKE
        },
      },
    },
  },
};
