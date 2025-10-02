package esepunittests

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

type GradeCalculator struct {
    grades []Grade
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{grades: make([]Grade, 0)}
}

func (gc *GradeCalculator) GetFinalGrade() string {
    numeric := gc.calculateNumericalGrade()
    switch {
    case numeric >= 90:
        return "A"
    case numeric >= 80:
        return "B"
    case numeric >= 70:
        return "C"
    case numeric >= 60:
        return "D"
    default:
        return "F"
    }
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
    gc.grades = append(gc.grades, Grade{Name: name, Grade: grade, Type: gradeType})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
    assignAvg := computeAverage(gc.grades, Assignment)
    examAvg := computeAverage(gc.grades, Exam)
    essayAvg := computeAverage(gc.grades, Essay)

    weighted := float64(assignAvg)*0.50 + float64(examAvg)*0.35 + float64(essayAvg)*0.15
    return int(weighted)
}

func computeAverage(grades []Grade, gradeType GradeType) int {
    count, sum := 0, 0
    for _, g := range grades {
        if g.Type == gradeType {
            sum += g.Grade
            count++
        }
    }
    if count == 0 {
        return 0
    }
    return sum / count
}