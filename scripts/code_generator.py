def generate_ld(dst):
    src = ["B","C","D","E","H","L","A"]

    for _src in src:
        function = "func (cpu *CPU) LD_"+dst.upper()+"_"+_src+"() int { \n return cpu.ld_dst_src(cpu.Registers."+dst.upper() + ", cpu.Registers."+_src+")\n}"
        print(function + "\n")


def generate_ld_dst_loc_src(dst, src):
    function = "func (cpu *CPU) LD_" + dst.upper() + "_LOC_" + src.upper() + "() int { \n return cpu.ld_dst_loc_src(cpu.Registers." + dst.upper() + ", cpu.Registers." + src.upper() + ")\n}"
    print(function + "\n")


def generate_register_d8():
    registers = ["B","D","H","C","A","E","C"]
    for register in registers:
        function = "func (cpu *CPU) LD_"+register + "_D8() int { \n return cpu.ld_dst_d8(cpu.Registers." + register + ")\n}"
        print(function)


def generate_inc_register():
    registers = ["B","D","H","C","E","L","A"]
    for register in registers:
        function = "func (cpu *CPU) INC_"+register+"() int { \n return cpu.inc_register(cpu.Registers." + register + ")\n}"
        print(function)

def generate_inc_register_6():
    registers = ["BC","DE","HL","SP"]
    for register in registers:
        function = "func (cpu *CPU) INC_"+register+"() int { \n return cpu.inc_register_sixteen(cpu.Registers." + register + ")\n}"
        print(function)

def generate_dec_register():
    registers = ["B","D","H","C","E","L","A"]
    for register in registers:
        function = "func (cpu *CPU) DEC_"+register+"() int { \n return cpu.dec_register(cpu.Registers." + register + ")\n}"
        print(function)

def generate_dec_register_6():
    registers = ["BC","DE","HL"]
    for register in registers:
        function = "func (cpu *CPU) DEC_"+register+"() int { \n return cpu.dec_register_sixteen(cpu.Registers." + register + ")\n}"
        print(function)


def generate_add_hl_sixteen():
    registers = ["BC","DE","HL","SP"]

    for register in registers:
        if register == "SP":
            r = "cpu.Registers.HL, cpu.SP"
        else:
            r = "cpu.Registers.HL, cpu.Registers." + register
        function = "func (cpu *CPU) ADD_HL_"+register+"() int { \n return cpu.add_dst_src_sixteen(" + r + ")\n}"
        print(function)

def generate_add_a_register():
    registers = ["B","C","D","E","H","L","A"]
    for register in registers:
        function = "func (cpu *CPU) ADD_A_"+register+"() int { \n return cpu.add_dst_src(cpu.Registers.A, cpu.Registers." + register + ".Get())\n}"
        print(function)


def generate_sub_a_register():
    registers = ["B","C","D","E","H","L","A"]
    for register in registers:
        function = "func (cpu *CPU) SUB_"+register+"() int { \n return cpu.sub_a_src(cpu.Registers." + register + ")\n}"
        print(function)


def generate_and_a_register():
    registers = ["B","C","D","E","H","L","A"]
    for register in registers:
        function = "func (cpu *CPU) AND_"+register+"() int { \n return cpu.and(cpu.Registers." + register + ")\n}"
        print(function)

def generate_or_a_register():
    registers = ["B","C","D","E","H","L","A"]
    for register in registers:
        function = "func (cpu *CPU) OR_"+register+"() int { \n return cpu.or(cpu.Registers." + register + ")\n}"
        print(function)

def generate_xor_a_register():
    registers = ["B","C","D","E","H","L"]
    for register in registers:
        function = "func (cpu *CPU) XOR_"+register+"() int { \n return cpu.xor(cpu.Registers." + register + ".Get())\n}"
        print(function)

def generate_cp():
    registers = ["B","C","D","E","H","L"]
    for register in registers:
        function = "func (cpu *CPU) CP_"+register+"() int { \n return cpu.cp(cpu.Registers." + register + ".Get())\n}"
        print(function)

def generate_pop():
    registers = ["DE","HL","AF"]
    for register in registers:
        function = "func (cpu *CPU) POP_"+register+"() int { \n return cpu.pop(cpu.Registers." + register + ")\n}"
        print(function)

def generate_push():
    registers = ["DE","HL","AF"]
    for register in registers:
        function = "func (cpu *CPU) PUSH_"+register+"() int { \n return cpu.push(cpu.Registers." + register + ")\n}"
        print(function)

def generate_adc():
    registers = ["B","C","D","E","H","L","A"]
    for register in registers:
        function = "func (cpu *CPU) ADC_A_"+register+"() int { \n return cpu.adc(cpu.Registers." + register + ".Get())\n}"
        print(function)

def generate_sdc():
    registers = ["B","C","D","E","H","L","A"]
    for register in registers:
        function = "func (cpu *CPU) SBC_A_"+register+"() int { \n return cpu.sbc(cpu.Registers." + register + ".Get())\n}"
        print(function)

def generate_rst():
    for x in range(8):
        function = "func (cpu *CPU) RST_" + str(x) + "() int { \n return cpu.rst(" + str(x) + ")\n}"
        print(function)

def generate_cb_rl():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        f = "return cpu.cb_rl8(cpu.Registers." + register +")"
        if register == "HL":
            f = "return cpu.cb_rl16(cpu.Registers." + register +")"
        function = "func (cpu *CPU) RL_" + register + "() int { \n" + f + "\n}"
        print(function)
def generate_cb_rr():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        size = "16" if register == "HL" else "8"
        f = f"return cpu.cb_rr{size}(cpu.Registers." + register +")"
        function = "func (cpu *CPU) RR_" + register + "() int { \n" + f + "\n}"
        print(function)


def generate_cb_rlc():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        size = "16" if register == "HL" else "8"
        f = f"return cpu.cb_rlc{size}(cpu.Registers." + register +")"
        function = "func (cpu *CPU) RLC_" + register + "() int { \n" + f + "\n}"
        print(function)

def generate_cb_rrc():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        size = "16" if register == "HL" else "8"
        f = f"return cpu.cb_rrc{size}(cpu.Registers." + register +")"
        function = "func (cpu *CPU) RRC_" + register + "() int { \n" + f + "\n}"
        print(function)

def generate_sla():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        size = "16" if register == "HL" else "8"
        f = f"return cpu.cb_sla{size}(cpu.Registers." + register + ")"
        function = "func (cpu *CPU) SLA_" + register + "() int { \n" + f + "\n}"
        print(function)

def generate_sra():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        size = "16" if register == "HL" else "8"
        f = f"return cpu.cb_sra{size}(cpu.Registers." + register + ")"
        function = "func (cpu *CPU) SRA_" + register + "() int { \n" + f + "\n}"
        print(function)

def generate_swap():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        f = f"return cpu.cb_swap(cpu.Registers." + register + ")"
        function = "func (cpu *CPU) SWAP_" + register + "() int { \n" + f + "\n}"
        print(function)

def generate_srl():
    registers = ["B","C","D","E","H","L","HL","A"]

    for register in registers:
        f = f"return cpu.cb_srl(cpu.Registers." + register + ")"
        function = "func (cpu *CPU) SRL_" + register + "() int { \n" + f + "\n}"
        print(function)

def generate_bit():
    registers = ["B","C","D","E","H","L","HL","A"]
    for x in range(8):
        for register in registers:
            f = f"return cpu.cb_bit(" + str(x) + ", cpu.Registers." + register + ")"

            function = "func (cpu *CPU) BIT_" + str(x) + "_" +  register + "() int { \n" + f + "\n}\n"
            print(function)

def generate_res():
    registers = ["B","C","D","E","H","L","HL","A"]
    for x in range(8):
        for register in registers:
            f = f"return cpu.cb_res_set(" + str(x) + ", 0, cpu.Registers." + register + ")"

            function = "func (cpu *CPU) RES_" + str(x) + "_" +  register + "() int { \n" + f + "\n}\n"
            print(function)

def generate_set():
    registers = ["B","C","D","E","H","L","HL","A"]
    for x in range(8):
        for register in registers:
            f = f"return cpu.cb_res_set(" + str(x) + ", 1, cpu.Registers." + register + ")"

            function = "func (cpu *CPU) SET_" + str(x) + "_" +  register + "() int { \n" + f + "\n}\n"
            print(function)


def generate_switch():
    for x in range(0x1d, 0xff+1):
        print("case " + hex(x) + ":\ncycles=cpu.INC_E()")

def generate_cb_switch(name, starting, length):
    registers = ["B","C","D","E","H","L","HL","A"]
    loc = starting
    count = 0
    index = 0
    for x in range(length):
        print("case " + hex(loc) + ":\ncycles=cpu." + name + "_" + str(index) + "_" + registers[count] + "()")
        count += 1

        if count == len(registers):
            count = 0
            index += 1

        loc += 1

def generate_eights(name, starting):
    registers = ["B","C","D","E","H","L","HL","A"]
    loc = starting
    for register in registers:
        print("case " + hex(loc) + ":\ncycles=cpu." + name + "_" + register + "()")
        loc += 1

generate_eights("RLC",0x0)
generate_eights("RRC",0x8)