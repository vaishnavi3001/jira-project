import { constant_values } from "../../constants";

describe("Checks if the project can be deleted", () => {

    it("should test if the project can be deleted", () => {
      localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)

      cy.visit(`${constant_values.PROJECT_LIST}`);
      cy.contains("sample_proj").click();
      cy.wait(2000);
      cy.contains("Delete Project").click();
      cy.wait(2000);
    //   cy.get("[id=email]").type("pypalkar23");
    //   cy.get("[id=password]").type("mandar");
  
      // cy.get("[id=login-button]").click();
    
    });
  
  });