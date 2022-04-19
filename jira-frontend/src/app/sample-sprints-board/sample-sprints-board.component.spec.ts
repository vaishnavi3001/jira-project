import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SampleSprintsBoardComponent } from './sample-sprints-board.component';

describe('SampleSprintsBoardComponent', () => {
  let component: SampleSprintsBoardComponent;
  let fixture: ComponentFixture<SampleSprintsBoardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SampleSprintsBoardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SampleSprintsBoardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
