// import { constant_values } from "../../constants";

// describe("Checks if the issue can be created", () => {

//     it("should test if the issue can be created", () => {
//       localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)
//       cy.visit('http://localhost:4200/home/sprint/1/details');
//       cy.visit('http://localhost:4200/home/1/issues/create');
      
//       cy.get("[id=title]").type("sample issue");
//       cy.get("[id=p_name]").type("sample project");
//       //cy.get("[id=issue_type]").click();

//       cy.get('select').select('Bug')

//       cy.contains('Bug').click();
//       cy.get("[id=sprint]").click();
      
//       cy.contains('Sprint 2').click();
//       cy.get("[id=desc]").type("pypalkar23");
//       cy.get("[id=assignee]").click();
//       cy.contains('Assignee 1').click();
//       cy.get("[id=reporter]").click();
//       cy.contains('Assignee 2').click();
//       cy.get("[id=create_issue]").click();
    
//     });
  
//   });