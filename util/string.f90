subroutine rubber(input, length)
    use json_module
    implicit none

    type(json_core) :: json
    type(json_value),pointer :: p
    type(json_value),POINTER :: scheduleJson
    INTEGER :: length
    INTEGER :: i
    CHARACTER(len=length) :: input
    logical :: found

    call json%create_object(p, '')
    call json%deserialize(p, input)

    call json%get(p, 'simulation.schedule', scheduleJson, found)

    call json%print(scheduleJson)

    do i = 1, json%count(scheduleJson)
        print *, i
    end do

end subroutine