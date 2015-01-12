package main

import (
	"fmt"
	"testing"

	. "github.com/wwitzel3/waynewitzel.com-code/util/patch"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type InterloveSuite struct {
	r Restorer
}

var _ = Suite(&InterloveSuite{})

func (s *InterloveSuite) SetUpTest(c *C) {
	testingProject := &Project{tag: "p1-testing"}
	testingEmail := &Email{tag: "e1-testing"}

	s.r = Patch(&doCall, func(svc Services) (interface{}, error) {
		if svc.project != nil {
			return testingProject, nil
		} else if svc.email != nil {
			return testingEmail, nil
		}
		return nil, fmt.Errorf("no matching testing service found")
	})
}

func (s *InterloveSuite) TearDownTest(c *C) {
	s.r()
}

func (s *InterloveSuite) TestGetProject(c *C) {
	project, err := getProject()
	c.Check(err, IsNil)
	c.Check(project.Tag(), Equals, "p1-testing")
}

func (s *InterloveSuite) TestGetEmail(c *C) {
	email, err := getEmail()
	c.Check(err, IsNil)
	c.Check(email.Tag(), Equals, "e1-testing")
}
