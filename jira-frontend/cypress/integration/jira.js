describe('My First Test', () => {
    it('Does not do much!', () => {
      expect(true).to.equal(true)
    })
})

describe('JIRA Main', () => {
    it('loads successfully', () => {
        cy.visit('http://localhost:4200')
    })
})

describe('Route redirection', () => {
    it('loads successfully', () => {
        cy.visit('http://localhost:4200/xyz')
    })
})