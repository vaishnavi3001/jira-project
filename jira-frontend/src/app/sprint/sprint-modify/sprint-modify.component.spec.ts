import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SprintModifyComponent } from './sprint-modify.component';

describe('SprintModifyComponent', () => {
  let component: SprintModifyComponent;
  let fixture: ComponentFixture<SprintModifyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SprintModifyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SprintModifyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
