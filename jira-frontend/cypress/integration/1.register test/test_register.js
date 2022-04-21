import { constant_values } from "../../constants";


describe("Checks registeration", () => {

    it("should test the resistration", () => {
      cy.visit(`${constant_values.REGISTER}`);
      cy.get("[id=user_name]").type("sample_user99");
      cy.get("[id=first_name]").type("sample_first");
      cy.get("[id=last_name]").type("sample_last");
      cy.get("[id=email]").type("sample_user2@email.com");
      cy.get("[id=password]").type("sample_password");
  
      cy.get("[id=register-submit]").click();
    
    });
  
  });