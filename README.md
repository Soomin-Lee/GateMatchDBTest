# GateMatchDBTest
공통 코드를 최대한 분리해낸다는 느낌으로 작업. \
이 모듈을 사용하는 Repository의 기능이 변경된다고 해서 이 모듈 자체가 변경되면 안된다. \
인사정보 등 업데이트에 따라 달라질 수 있는 부분은 각각의 Repository에서 따로 담당한다.

- 특징점 DB CRUD (Name - FeatureBlob, 어느 배포처에 나가도 특징점 DB는 동일한 형식으로 관리될 것)
- Inference 및 Matching (Parallel Matching 기능 포함)
- DB 연결 (DatabaseManager)
