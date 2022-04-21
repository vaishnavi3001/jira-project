import { constant_values } from "../../constants";

describe("Checks if the user can be logged out.", () => {

    it("should test if the user can be logged out", () => {
      localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)
      cy.wait(2000);
      cy.contains("Logout").click();
      cy.wait(2000);

    });
  
  });