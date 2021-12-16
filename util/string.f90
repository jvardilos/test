subroutine rubber(input, length)
    use iso_c_binding
    use json_module
    implicit none

    
    type(json_file) :: json
    INTEGER :: length
    CHARACTER :: input(length)


    call json%initialize()

    call json%deserialize('{\"hello\":\"world\"}')
    call json%print()


    
    print *, input(1:length)

end subroutine