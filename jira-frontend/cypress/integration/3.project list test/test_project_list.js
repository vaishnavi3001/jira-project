import { constant_values } from "../../constants";

describe("Checks project list view", () => {

    it("should test the project list view page", () => {
      localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)

      cy.visit(`${constant_values.PROJECT_LIST}`);
      cy.wait(2000);
    //   cy.get("[id=email]").type("pypalkar23");
    //   cy.get("[id=password]").type("mandar");
  
      // cy.get("[id=login-button]").click();
    
    });
  
  });